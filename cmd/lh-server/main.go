package main

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/audit/lighthouse"
	"github.com/xwp/go-tide/src/audit/tide"
	"flag"
	"fmt"
	tideApi "github.com/xwp/go-tide/src/tide"
	"github.com/xwp/go-tide/src/tide/api"
	"github.com/xwp/go-tide/src/message/sqs"
	"time"
	"log"
	"github.com/xwp/go-tide/src/env"
	"strconv"
	"github.com/xwp/go-tide/src/audit/ingest"
	"github.com/xwp/go-tide/src/audit/info"
	"github.com/xwp/go-tide/src/storage"
	"github.com/xwp/go-tide/src/storage/s3"
)

var (
	Version string // Set during build.
	Build   string // Set during build.

	// Use the interface so that we can test.
	TideClient tideApi.ClientInterface

	// Number of concurrent audits.
	bufferSize, _ = strconv.Atoi(env.GetEnv("LH_CONCURRENT_AUDITS", "5"))

	// Temp folder for downloaded files.
	tempFolder = env.GetEnv("LH_TEMP_FOLDER", "/tmp")

	// Tide API configuration.
	tideConfig = struct {
		id           string
		secret       string
		authEndpoint string
	}{
		env.GetEnv("TIDE_API_KEY", ""),
		env.GetEnv("TIDE_API_SECRET", ""),
		env.GetEnv("TIDE_API_AUTH_URL", ""),
	}

	// Lighthouse SQS configuration.
	lhConfig = struct {
		region string
		key    string
		secret string
		queue  string
	}{
		env.GetEnv("AWS_SQS_LH_REGION", ""),
		env.GetEnv("AWS_SQS_LH_KEY", ""),
		env.GetEnv("AWS_SQS_LH_SECRET", ""),
		env.GetEnv("AWS_SQS_LH_QUEUE_NAME", ""),
	}

	// Lighthouse SQS configuration.
	lhS3Config = struct {
		region string
		key    string
		secret string
		bucket string
	}{
		env.GetEnv("AWS_S3_LH_REGION", ""),
		env.GetEnv("AWS_S3_LH_KEY", ""),
		env.GetEnv("AWS_S3_LH_SECRET", ""),
		env.GetEnv("AWS_S3_LH_BUCKET_NAME", ""),
	}

	messageProvider message.MessageProvider
	storageProvider storage.StorageProvider
)

func main() {

	// Is the -version flag being used?
	bVersion := flag.Bool("version", false, "a bool")

	// Parse all flags.
	flag.Parse()

	// If -version is requested then echo the version details.
	if *bVersion {
		fmt.Printf("Version: %s\nBuild: %s\n", Version, Build)
	}

	// Prepare the Tide Client.
	TideClient = &api.Client{}
	err := TideClient.Authenticate(tideConfig.id, tideConfig.secret, tideConfig.authEndpoint)
	if err != nil {
		log.Fatal("Tide API Error:", err)
	}

	// Initiate a new Message provider.
	messageProvider = sqs.NewSqsProvider(lhConfig.region, lhConfig.key, lhConfig.secret, lhConfig.queue)

	// Initiate a new Storage provider.
	storageProvider = s3.NewS3Provider(lhS3Config.region, lhS3Config.key, lhS3Config.secret, lhS3Config.bucket)

	// Create a buffer for the amount of concurrent audits.
	buffer := make(chan struct{}, bufferSize)

	// Create a channel that receives messages from a queue.
	cMessage := messageChannel(messageProvider, buffer)

	// Poll the message channel until the program is forcefully exited.
	for {
		select {
		// Message received from the queue.
		case msg := <-cMessage:
			// Process the message in a go routine.
			go processMessage(msg, TideClient, buffer)
		}
	}
}

// messageChannel returns a channel of messages to be processed. The message provider gets polled for
// the next message and upon success it gets added to the channel.
func messageChannel(provider message.MessageProvider, buffer chan struct{}) <-chan *message.Message {

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
						log.Fatal(pErr)
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

	// @todo Provide better information.
	log.Println("Processing...")

	// An slice of processes that need to be performed on the message.
	// A slice ensures that they happen in the correct order.
	processes := []audit.Processor{
		&ingest.Processor{},
		&info.Processor{},
		&lighthouse.Processor{},
		&tide.Processor{},
	}

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
		}
	}

	// Remove message on success.
	if len(errors) == 0 {
		messageProvider.DeleteMessage(msg.ExternalRef)
	}

	// Release item from buffer so that next item can be polled.
	<-buffer
}
