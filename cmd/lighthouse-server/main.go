package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wptide/pkg/env"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/message/firestore"
	"github.com/wptide/pkg/message/mongo"
	"github.com/wptide/pkg/message/sqs"
	"github.com/wptide/pkg/payload"
	"github.com/wptide/pkg/process"
	"github.com/wptide/pkg/source"
	"github.com/wptide/pkg/storage"
	"github.com/wptide/pkg/storage/gcs"
	"github.com/wptide/pkg/storage/local"
	"github.com/wptide/pkg/storage/s3"
	tideApi "github.com/wptide/pkg/tide"
	"github.com/wptide/pkg/tide/api"
)

type processConfig struct {
	igTempFolder      string
	lhTempFolder      string
	lhStorageProvider storage.Provider
	resPayloaders     map[string]payload.Payloader
}

var (
	// Version is the binary version set during build.
	Version string

	// Build is the binary build number set during build.
	Build string

	// Lighthouse service configuration (as var so that we can also override).
	serviceConfig = getServiceConfig()

	// TideClient represents the instance that will write to the Tide API.
	TideClient tideApi.ClientInterface = &api.Client{}

	// Specify the payloads to be used for this service.
	payloaders = map[string]payload.Payloader{
		"tide": payload.TidePayload{
			Client: TideClient,
		},
		"local-file": payload.FilePayload{
			TerminateChannel: terminateChannel,
		},
	}

	// messageProvider is the primary source of messages to process.
	messageProvider = getMessageProvider(serviceConfig)

	/** Processes Configuration **/
	procCfg = processConfig{
		serviceConfig["app"]["temp_folder"],
		serviceConfig["app"]["temp_folder"],
		getStorageProvider(serviceConfig),
		payloaders,
	}

	/** Channels **/
	// Number of concurrent audits.
	//bufferSize, _ = strconv.Atoi(env.GetEnv("LH_CONCURRENT_AUDITS", "5"))

	// Create a channel that receives messages from a queue.
	cMessage = make(chan message.Message)

	// Create a channel that can terminate the app.
	terminateChannel = make(chan struct{}, 1)

	/** Flags **/
	bParseFlags = true

	// Display versions?
	bVersion = &[]bool{false}[0]

	// A single URL to process. Will not poll queue.
	flagURL = &[]string{""}[0]

	// Project visibility for URL. "private" or "public".
	flagVisibility = &[]string{"private"}[0]

	// The client login id in Tide API.
	flagClient = &[]string{"wporg"}[0]

	// Send to Stdout rather than API?
	flagOutput = &[]string{""}[0]

	// Process Functions.
	doProcess = executeProcessFunc
)

func main() {

	if messageProvider != nil {
		defer messageProvider.Close()
	}

	log.Println("Starting Lighthouse audit server.")

	if bParseFlags {
		// Is the -version flag being used?
		bVersion = flag.Bool("version", false, "print program version details")

		// An -output file to write results to. Will not send to API.
		flagOutput = flag.String("output", "", "send results to output file (json format)")

		// A -url to run a single audit. Will not poll a queue.
		flagURL = flag.String("url", "", "audit single message from url")

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
		conf := serviceConfig["tide"]
		err := TideClient.Authenticate(conf["key"], conf["secret"], conf["auth"])
		if err != nil {
			log.Println("Tide API Error:", err)
			terminateChannel <- struct{}{}
		}
	}

	// If -url is set then send a message using the URL.
	if *flagURL != "" {

		singleMessageType := "tide"
		conf := serviceConfig["tide"]
		singleMessageEndpoint := fmt.Sprintf("%s://%s/api/tide/%s/audit", conf["protocol"], conf["host"], conf["version"])

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
				SourceURL:           *flagURL,
				SourceType:          source.GetKind(*flagURL),
				Force:               true,
				Visibility:          *flagVisibility,
				RequestClient:       *flagClient,
				Audits: []*message.Audit{
					{
						Type:    "lighthouse",
						Options: &message.AuditOption{},
					},
				},
			}
		}()
	}

	// Start polling the messageProvider if we didn't provide a url.
	// Other processes can also queue the channel.
	if *flagURL == "" {
		log.Println("Polling message provider.")
		pollProvider(cMessage, messageProvider)
	}

	// Poll the message channel until the program is forcefully exited.
	for {
		select {
		case msg := <-cMessage:
			// Process Message.
			wg := sync.WaitGroup{}
			wg.Add(1)
			err := processMessage(msg, &wg)
			if err != nil {
				fmt.Println(err)
			}
			wg.Wait()

		case <-terminateChannel:
			// Terminate signal received.
			goto terminated
		}
	}

terminated:
	fmt.Println("Server terminated.")
}

// processMessage sends the message through a sequential process.
func processMessage(msg message.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Initiate results pointer.
	results := &process.Result{}

	// Ingest.
	ingestProc := &process.Ingest{
		TempFolder: procCfg.igTempFolder,
	}
	if err := doProcess(ingestProc, msg, results); err != nil {
		return err
	}

	// Info.
	infoProc := &process.Info{}
	if err := doProcess(infoProc, msg, results); err != nil {
		return err
	}

	// Lighthouse.
	lighthouseProc := &process.Lighthouse{
		TempFolder:      procCfg.lhTempFolder,
		StorageProvider: procCfg.lhStorageProvider,
	}
	if err := doProcess(lighthouseProc, msg, results); err != nil {
		return err
	}

	// Prepare Response.
	responseProc := &process.Response{
		Payloaders: payloaders,
	}
	if err := doProcess(responseProc, msg, results); err != nil {
		return err
	}

	// Clean up.
	res := *results
	if res["responseSuccess"] != nil && res["responseSuccess"].(bool) {
		// Output message.
		log.Println(res["responseMessage"])

		// Delete message from provider.
		err := messageProvider.DeleteMessage(msg.ExternalRef)
		if err != nil {
			return err
		}

		log.Println("'" + msg.Title + "' removed from message queue.")
	}

	return nil
}

// pollProvider polls the message provider for
// the next message and upon success it gets added to the channel.
func pollProvider(c chan message.Message, provider message.Provider) {

	// Run this concurrently.
	go func() {
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

				// Send message to channel.
				c <- *msg
			}
			time.Sleep(time.Second * time.Duration(2))
		}
	}()
}

// executeProcessFunc takes the incoming process and executes it.
// This is assigned to `doProcess` to make it testable.
func executeProcessFunc(proc process.Processor, msg message.Message, result *process.Result) error {
	proc.SetMessage(msg)
	proc.SetResults(result)
	return proc.Do()
}

// getStorageProvider returns a storage provider given the provided configurations from
// the environment variables.
func getStorageProvider(config map[string]map[string]string) storage.Provider {
	switch config["app"]["storage_provider"] {
	case "s3":
		conf := config["aws"]
		return s3.NewS3Provider(conf["s3_region"], conf["key"], conf["secret"], conf["s3_bucket"])
	case "gcs":
		conf := config["gcp"]
		return gcs.NewCloudStorageProvider(context.Background(), conf["project"], conf["gcs_bucket"])
	case "local":
		return local.NewLocalStorage(config["app"]["server_path"], config["app"]["local_path"])
	default:
		return nil
	}
}

// getStorageProvider returns a message/queue provider given the provided configurations
// from the environment variables.
func getMessageProvider(config map[string]map[string]string) message.Provider {

	switch config["app"]["message_provider"] {
	case "sqs":
		conf := config["aws"]
		return sqs.NewSqsProvider(conf["sqs_region"], conf["key"], conf["secret"], conf["sqs_queue"])
	case "firestore":
		conf := config["gcp"]
		fp, _ := firestore.New(context.Background(), conf["project"], conf["gcf_queue"])
		return fp
	case "local":
		conf := config["mongo"]
		mp, _ := mongo.New(context.Background(), conf["user"], conf["pass"], conf["host"], conf["database"], conf["queue"], nil)
		return mp
	default:
		return nil
	}

}

func getServiceConfig() map[string]map[string]string {
	return map[string]map[string]string{
		"app": {
			"storage_provider": env.GetEnv("LH_STORAGE_PROVIDER", "local"),
			"message_provider": env.GetEnv("LH_MESSAGE_PROVIDER", "local"),
			"temp_folder":      env.GetEnv("LH_TEMP_FOLDER", "/tmp"),
			"server_path":      "/srv/data",
			"local_path":       "lighthouse",
		},
		"aws": {
			"key":        env.GetEnv("AWS_API_KEY", ""),
			"secret":     env.GetEnv("AWS_API_SECRET", ""),
			"sqs_region": env.GetEnv("AWS_SQS_REGION", ""),
			"sqs_queue":  env.GetEnv("AWS_SQS_QUEUE_LH", ""),
			"s3_region":  env.GetEnv("AWS_S3_REGION", ""),
			"s3_bucket":  env.GetEnv("AWS_S3_BUCKET_NAME", ""),
		},
		"gcp": {
			"project":    env.GetEnv("GCP_PROJECT", ""),
			"gcs_bucket": env.GetEnv("GCS_BUCKET_NAME", ""),
			"gcf_queue":  env.GetEnv("GCF_QUEUE_LH", "queue-lighthouse"),
		},
		"mongo": {
			"user":     env.GetEnv("MONGO_DATABASE_USERNAME", ""),
			"pass":     env.GetEnv("MONGO_DATABASE_PASSWORD", ""),
			"host":     "localhost:27017",
			"database": env.GetEnv("MONGO_DATABASE_NAME", "queue"),
			"queue":    env.GetEnv("MONGO_QUEUE_LH", "lighthouse"),
		},
		"tide": {
			"key":      env.GetEnv("API_KEY", ""),
			"secret":   env.GetEnv("API_SECRET", ""),
			"auth":     env.GetEnv("API_AUTH_URL", ""),
			"host":     env.GetEnv("API_HTTP_HOST", ""),
			"protocol": env.GetEnv("API_PROTOCOL", ""),
			"version":  env.GetEnv("API_VERSION", ""),
		},
	}
}
