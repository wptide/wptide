package main

import (
	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/audit/lighthouse"
	"github.com/wptide/pkg/audit/tide"
	tideApi "github.com/wptide/pkg/tide"
	"github.com/wptide/pkg/tide/api"
	"github.com/wptide/pkg/message/sqs"
	"time"
	"log"
	tidelog "github.com/wptide/pkg/log"
	"github.com/wptide/pkg/env"
	"strconv"
	"github.com/wptide/pkg/audit/ingest"
	"github.com/wptide/pkg/audit/info"
	"github.com/wptide/pkg/storage"
	"github.com/wptide/pkg/storage/s3"
	"flag"
	"github.com/wptide/pkg/source"
	"fmt"
	"os"
)

var (
	/** Compile time variables. **/
	Version string // Set during build.
	Build   string // Set during build.

	/** Initialize interface instances **/
	// Use the interface so that we can test.
	TideClient tideApi.ClientInterface = &api.Client{}

	// messageProvider is the primary source of messages to process.
	messageProvider message.MessageProvider = sqs.NewSqsProvider(lhConfig.region, lhConfig.key, lhConfig.secret, lhConfig.queue)

	// storageProvider is the cloud storage service for storing files.
	storageProvider storage.StorageProvider = s3.NewS3Provider(lhS3Config.region, lhS3Config.key, lhS3Config.secret, lhS3Config.bucket)

	// Temp folder for downloaded files.
	tempFolder = env.GetEnv("LH_TEMP_FOLDER", "/tmp")

	/** Service Configurations **/
	// Tide API configuration.
	tideConfig = struct {
		id           string
		secret       string
		authEndpoint string
		host         string
		protocol     string
		version      string
	}{
		env.GetEnv("TIDE_API_KEY", ""),
		env.GetEnv("TIDE_API_SECRET", ""),
		env.GetEnv("TIDE_API_AUTH_URL", ""),
		env.GetEnv("TIDE_API_HOST", ""),
		env.GetEnv("TIDE_API_PROTOCOL", ""),
		env.GetEnv("TIDE_API_VERSION", ""),
	}

	// Lighthouse SQS configuration.
	lhConfig = struct {
		region string
		key    string
		secret string
		queue  string
	}{
		env.GetEnv("LH_SQS_REGION", ""),
		env.GetEnv("LH_SQS_KEY", ""),
		env.GetEnv("LH_SQS_SECRET", ""),
		env.GetEnv("LH_SQS_QUEUE_NAME", ""),
	}

	// Lighthouse SQS configuration.
	lhS3Config = struct {
		region string
		key    string
		secret string
		bucket string
	}{
		env.GetEnv("LH_S3_REGION", ""),
		env.GetEnv("LH_S3_KEY", ""),
		env.GetEnv("LH_S3_SECRET", ""),
		env.GetEnv("LH_S3_BUCKET_NAME", ""),
	}

	/** Channels **/
	// Number of concurrent audits.
	bufferSize, _ = strconv.Atoi(env.GetEnv("LH_CONCURRENT_AUDITS", "5"))
	buffer        = make(chan struct{}, bufferSize)

	// Create a channel that can terminate the app.
	terminateChannel = make(chan struct{}, 1)

	// Create a channel that receives messages from a queue.
	cMessage chan *message.Message

	/** Flags **/
	bParseFlags = true

	// Display versions?
	bVersion = &[]bool{false}[0]

	// A single URL to process. Will not poll queue.
	flagUrl = &[]string{""}[0]

	// Project visibility for URL. "private" or "public"
	flagVisibility = &[]string{"private"}[0]

	// The client login id in Tide API.
	flagClient = &[]string{"wporg"}[0]

	// Send to Stdout rather than API?
	flagOutput = &[]string{""}[0]

	/** Processes **/
	// A slice of processes (in order) that need to be performed on a message.
	processes = []audit.Processor{
		&ingest.Processor{},
		&info.Processor{},
		&lighthouse.Processor{},
		&tide.Processor{},
	}
)

func main() {

	if bParseFlags {
		// Is the -version flag being used?
		bVersion = flag.Bool("version", false, "print program version details")

		// An -output file to write results to. Will not send to API.
		flagOutput = flag.String("output", "", "send results to output file (json format)")

		// A -url to run a single audit. Will not poll a queue.
		flagUrl = flag.String("url", "", "audit single message from url")

		// A -visibility to run a single audit. Will not poll a queue.
		flagVisibility = flag.String("visibility", "public", `"private" or "public" - default "pubic"`)

		// The -client login name in Tide API.
		flagClient = flag.String("client", "wporg", `Tide API client to attribute project to - default "wporg"`)

		// Parse all flags.
		flag.Parse()
	}

	// If -version is requested then echo the version details.
	if *bVersion {
		oldFlags := log.Flags()
		log.SetFlags(0)
		log.Printf("Version: %s\nBuild: %s\n", Version, Build)
		log.SetFlags(oldFlags)
	}

	// If -url is set then prepare cMessage without a provider.
	if *flagUrl != "" {
		cMessage = make(chan *message.Message, 1)
		cMessage <- &message.Message{
			ResponseAPIEndpoint: fmt.Sprintf("%s://%s/api/tide/%s/audit", tideConfig.protocol, tideConfig.host, tideConfig.version),
			Title:               "Single Project",
			Content:             "Single project URL provided.",
			SourceURL:           *flagUrl,
			SourceType:          source.GetKind(*flagUrl),
			Force:               true,
			Visibility:          *flagVisibility,
			RequestClient:       *flagClient,
		}
	}

	// If -output is set add it to Tide process OutputFile.
	if *flagOutput != "" {
		var procs []audit.Processor

		for _, proc := range processes {
			if proc.Kind() != "tide" {
				procs = append(procs, proc)
			} else {
				procs = append(procs, &tide.Processor{OutputFile: *flagOutput})
			}
		}

		processes = procs
	}

	// Prepare the Tide Client.
	err := TideClient.Authenticate(tideConfig.id, tideConfig.secret, tideConfig.authEndpoint)
	if err != nil {
		log.Println("Tide API Error:", err)
		terminateChannel <- struct{}{}
	}

	// Create a channel that receives messages from a queue, if it does not yet exist.
	if cMessage == nil {
		cMessage = messageChannel(messageProvider, buffer)
	}

	// Poll the message channel until the program is forcefully exited.
	for {
		select {
		// Terminate signal received.
		case <-terminateChannel:
			break;
			// Message received from the queue.
		case msg := <-cMessage:
			// If its a single process then don't run as goroutine.
			if *flagUrl != "" {
				// We still need to prime the bugger for processMessage() to be able to exit.
				buffer <- struct{}{}
				processMessage(msg, TideClient, buffer)
				os.Exit(0)
			} else {
				// Process the message in a go routine.
				go processMessage(msg, TideClient, buffer)
			}
		}
	}
}

// messageChannel returns a channel of messages to be processed. The message provider gets polled for
// the next message and upon success it gets added to the channel.
func messageChannel(provider message.MessageProvider, buffer chan struct{}) chan *message.Message {

	// Create the message channel.
	c := make(chan *message.Message)

	// Run this concurrently.
	go func(b chan struct{}) {
		for {

			// Get message from provider.
			msg, err := provider.GetNextMessage()

			// Handle provider errors.
			if err != nil {
				// If its a Provider error we might need to panic and fail.
				if pErr, ok := err.(*message.ProviderError); ok {
					switch pErr.Type {
					case message.ErrCritcal:
						log.Println(pErr)
						terminateChannel <- struct{}{}
						break
					case message.ErrOverQuota:
						log.Println(pErr, "delaying for 60 seconds")
						time.Sleep(time.Second * time.Duration(60))
					}
				}
			}

			log.Println("Polling...")

			// If message has been retrieved add it to the channel.
			if msg != nil {
				// Block if the buffer is full.
				b <- struct{}{}

				// Send message to channel.
				c <- msg
			}
			time.Sleep(time.Second * time.Duration(2))
		}
	}(buffer)

	// Return the message channel.
	return c
}

// processMessage takes an SQS message and runs it through an audit process.
func processMessage(msg *message.Message, client tideApi.ClientInterface, buffer <-chan struct{}) {

	tidelog.Log(msg.Title, "Processing...")

	var errors []error

	// Initialise result with:
	// - Tide client reference,
	// - Temp folder for downloaded/extracted files,
	// - Audits to run
	result := &audit.Result{
		"client":     &client,
		"tempFolder": tempFolder,
		"audits": []string{
			"lighthouse",
		},
		"fileStore": &storageProvider,
	}

	// Pointer that we can use directly.
	r := *result

	// Run through each processor and update the result.
	// Note: The result is a pointer so we're passing by reference.
	for _, proc := range processes {
		proc.Process(*msg, result)
		kind := proc.Kind()
		if r[kind+"Error"] != nil {
			errors = append(errors, r[kind+"Error"].(error))
			tidelog.Log(msg.Title, fmt.Sprintf("%s process error: %s", kind, r[kind+"Error"].(error)))
		}
	}

	// Remove message on success.
	if len(errors) == 0 {
		messageProvider.DeleteMessage(msg.ExternalRef)
		tidelog.Log(msg.Title, "Completed without errors.")
	} else {
		tidelog.Log(msg.Title, "Completed with some errors.")
	}

	// Release item from buffer so that next item can be polled.
	<-buffer
}
