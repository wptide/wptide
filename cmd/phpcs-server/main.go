package main

import (
	"github.com/wptide/pkg/message"
	tideApi "github.com/wptide/pkg/tide"
	"github.com/wptide/pkg/tide/api"
	"github.com/wptide/pkg/message/sqs"
	"time"
	"log"
	"github.com/wptide/pkg/env"
	"strconv"
	"github.com/wptide/pkg/storage"
	"github.com/wptide/pkg/storage/s3"
	"flag"
	"github.com/wptide/pkg/process"
	"github.com/wptide/pkg/payload"
	"errors"
	"fmt"
	"github.com/wptide/pkg/source"
	"github.com/wptide/pkg/pipe"
	commonProcess "github.com/xwp/go-tide/src/process"
	commonPayload "github.com/xwp/go-tide/src/payload"
)

type processConfig struct {
	igTempFolder         string
	phpcsTempFolder      string
	phpcsStorageProvider storage.StorageProvider
	resPayloaders        map[string]payload.Payloader
}

var (
	/** Compile time variables. **/
	Version string // Set during build.
	Build   string // Set during build.

	/** Initialize interface instances **/
	// Use the interface so that we can test.
	TideClient tideApi.ClientInterface = &api.Client{}

	// Specify the payloads to be used for this service.
	payloaders = map[string]payload.Payloader{
		"tide": payload.TidePayload{
			Client: TideClient,
		},
		"local-file": commonPayload.FilePayload{
			TerminateChannel: terminateChannel,
		},
	}

	// messageProvider is the primary source of messages to process.
	messageProvider message.MessageProvider = sqs.NewSqsProvider(queueConfig.region, queueConfig.key, queueConfig.secret, queueConfig.queue)

	// storageProvider is the cloud storage service for storing files.
	storageProvider storage.StorageProvider = s3.NewS3Provider(s3Config.region, s3Config.key, s3Config.secret, s3Config.bucket)

	// Temp folder for downloaded files.
	tempFolder = env.GetEnv("PHPCS_TEMP_FOLDER", "/tmp")

	/** Processes Configuration **/
	procCfg = &processConfig{
		tempFolder,
		tempFolder,
		storageProvider,
		payloaders,
	}

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
		env.GetEnv("API_KEY", ""),
		env.GetEnv("API_SECRET", ""),
		env.GetEnv("API_AUTH_URL", ""),
		env.GetEnv("API_HTTP_HOST", ""),
		env.GetEnv("API_PROTOCOL", ""),
		env.GetEnv("API_VERSION", ""),
	}

	// SQS configuration.
	queueConfig = struct {
		region string
		key    string
		secret string
		queue  string
	}{
		env.GetEnv("AWS_SQS_REGION", ""),
		env.GetEnv("AWS_API_KEY", ""),
		env.GetEnv("AWS_API_SECRET", ""),
		env.GetEnv("AWS_SQS_QUEUE_PHPCS", ""),
	}

	// Lighthouse SQS configuration.
	s3Config = struct {
		region string
		key    string
		secret string
		bucket string
	}{
		env.GetEnv("AWS_S3_REGION", ""),
		env.GetEnv("AWS_API_KEY", ""),
		env.GetEnv("AWS_API_SECRET", ""),
		env.GetEnv("AWS_S3_BUCKET_NAME", ""),
	}

	/** Channels **/
	// Number of concurrent audits.
	bufferSize, _ = strconv.Atoi(env.GetEnv("PHPCS_CONCURRENT_AUDITS", "5"))
	buffer        = make(chan struct{}, bufferSize)

	// Create a channel that can terminate the app.
	terminateChannel = make(chan struct{}, 1)

	// Create a channel that receives messages from a queue.
	cMessage = make(chan message.Message)

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
)

// initProcesses creates a number of processes for the pipeline.
// With each process we connect the "In" and "Out" channels to subsequent processes.
func initProcesses(source <-chan message.Message, config *processConfig) ([]process.Processor, error) {

	// Make sure we have a valid channel.
	if source == nil {
		return nil, errors.New("no valid source channel")
	}

	// Make sure we have our config.
	if config == nil {
		return nil, errors.New("no config provided")
	}
	if config.igTempFolder == "" {
		return nil, errors.New("no temp folder for ingest process")
	}
	if config.phpcsTempFolder == "" {
		return nil, errors.New("no temp folder for phpcs process")
	}
	if config.phpcsStorageProvider == nil {
		return nil, errors.New("no storage provider for phpcs process")
	}
	if config.resPayloaders == nil {
		return nil, errors.New("no response payloaders to send to")
	}

	// Create and connect processes.
	ingest := &process.Ingest{
		In:         source,
		Out:        make(chan process.Processor),
		TempFolder: config.igTempFolder,
	}

	info := &process.Info{
		In:  ingest.Out,
		Out: make(chan process.Processor),
	}

	intercept := &Intercept{
		In:  info.Out,
		Out: make(chan process.Processor),
	}

	phpcs := &process.Phpcs{
		In:         intercept.Out,
		Out:        make(chan process.Processor),
		TempFolder: config.phpcsTempFolder,
		Config: map[string]interface{}{
			"parallel": 1,
		},
		StorageProvider: config.phpcsStorageProvider,
	}

	// Not chaining response, so no need for Out channel.
	response := &process.Response{
		In:         phpcs.Out,
		Out:        make(chan process.Processor),
		Payloaders: config.resPayloaders,
	}

	sink := &commonProcess.Sink{
		In: response.Out,
		MessageProvider: messageProvider,
	}

	return []process.Processor{
		ingest,
		info,
		intercept,
		phpcs,
		response,
		sink,
	}, nil
}

func main() {

	log.Println("Starting PHPCS audit server.")

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

	// Prepare the Tide Client if we're not writing to file.
	if *flagOutput == "" {
		log.Println("Authenticating with Tide API.")
		err := TideClient.Authenticate(tideConfig.id, tideConfig.secret, tideConfig.authEndpoint)
		if err != nil {
			log.Println("Tide API Error:", err)
			terminateChannel <- struct{}{}
		}
	}

	/** Processes **/
	log.Println("Initializing processes.")
	processes, _ := initProcesses(
		cMessage,
		procCfg,
	)

	// Create and run the pipe.
	go func() {
		// Create New Pipe with processes.
		log.Println("Starting processing pipeline.")
		pipeline := pipe.WithProcesses(processes...)

		err := pipeline.Run()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// If -url is set then send a message using the URL.
	if *flagUrl != "" {

		singleMessageType := "tide"
		singleMessageEndpoint := fmt.Sprintf("%s://%s/api/tide/%s/audit", tideConfig.protocol, tideConfig.host, tideConfig.version)

		// If -output is set add it to Tide process OutputFile.
		if *flagOutput != "" {
			singleMessageType = "local-file"
			singleMessageEndpoint = *flagOutput
		}

		go func() {
			cMessage <- message.Message{
				ResponseAPIEndpoint: singleMessageEndpoint,
				Title:               "Single Project",
				Content:             "Single project URL provided.",
				PayloadType:         singleMessageType,
				SourceURL:           *flagUrl,
				SourceType:          source.GetKind(*flagUrl),
				Force:               true,
				Visibility:          *flagVisibility,
				RequestClient:       *flagClient,
				Audits: &[]message.Audit{
					{
						Type: "phpcs",
						Options: &message.AuditOption{
							Standard: "wordpress",
						},
					},
					{
						Type: "phpcs",
						Options: &message.AuditOption{
							Standard: "phpcompatibility",
						},
					},
				},
			}
		}()
	}

	// Start polling the messageProvider if we didn't provide a url.
	// Other processes can also queue the channel.
	if *flagUrl == "" {
		log.Println("Polling message provider.")
		pollProvider(cMessage, messageProvider, buffer)
	}

	// Poll the message channel until the program is forcefully exited.
	for {
		select {
		// Terminate signal received.
		case <-terminateChannel:
			goto terminated
		}
	}

terminated:
	fmt.Println("Server terminated.")
}

// pollProvider polls the message provider for
// the next message and upon success it gets added to the channel.
func pollProvider(c chan message.Message, provider message.MessageProvider, buffer chan struct{}) {

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

				log.Println("Dispatching '" + msg.Title + "'")

				// Block if the buffer is full.
				b <- struct{}{}

				// Send message to channel.
				c <- *msg
			}
			time.Sleep(time.Second * time.Duration(2))
		}
	}(buffer)
}
