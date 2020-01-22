package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/message/firestore"
	"github.com/wptide/pkg/message/mongo"
	"github.com/wptide/pkg/message/sqs"
	"github.com/wptide/pkg/storage/gcs"
	"github.com/wptide/pkg/storage/local"
	"github.com/wptide/pkg/storage/s3"
)

var (
	envTest = map[string]string{
		// Tide API config
		"API_AUTH_URL":  "http://tide.local/api/tide/v1/auth",
		"API_HTTP_HOST": "tide.local",
		"API_PROTOCOL":  "http",
		"API_KEY":       "tideapikey",
		"API_SECRET":    "tideapisecret",
		"API_VERSION":   "v1",
		//
		// AWS API settings
		"AWS_API_KEY":    "sqskey",
		"AWS_API_SECRET": "sqssecret",
		//
		// AWS S3 settings
		"AWS_S3_BUCKET_NAME": "test-bucket",
		"AWS_S3_REGION":      "us-west-2",
		//
		// AWS SQS settings
		"AWS_SQS_QUEUE_LH": "test-queue",
		"AWS_SQS_REGION":   "us-west-2",
		"AWS_SQS_VERSION":  "2012-11-05",
		//
		// LH Server settings
		"LH_CONCURRENT_AUDITS": "1",
	}
)

func Test_main(t *testing.T) {

	b := bytes.Buffer{}
	log.SetOutput(&b)

	var value string
	for key, val := range envTest {

		// Key is set, so retain the value when the test is finished.
		if value = os.Getenv(key); value != "" {
			os.Unsetenv(key)
			defer func() { os.Setenv(key, value) }()
		}

		// Set the test value.
		os.Setenv(key, val)
	}

	resetServiceConfig()

	// Use the mockTide for Tide
	currentClient := TideClient
	TideClient = &mockTide{}
	defer func() { TideClient = currentClient }()

	cMessageProvider := messageProvider
	messageProvider = &mockMessageProvider{}
	defer func() { messageProvider = cMessageProvider }()

	type args struct {
		messageChannel chan message.Message
		timeOut        time.Duration
		msg            message.Message
		parseFlags     bool
		version        bool
		authError      bool
		flagURL        *string
		flagOutput     *string
		flagVisibility *string
		altConfig      *processConfig
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"Run Main - Process Config Error (missing)",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				msg:            message.Message{},
				altConfig:      &processConfig{},
			},
		},
		{
			"Run Main",
			args{
				timeOut: 1,
				version: true,
			},
		},
		{
			"Run Main - Custom Message",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				msg:            message.Message{},
			},
		},
		{
			"Run Main - Version flag set",
			args{
				timeOut:    1,
				parseFlags: true,
			},
		},
		{
			"Run Main - Auth Error",
			args{
				timeOut:   1,
				authError: true,
			},
		},
		{
			"Run Main - Output Flag set",
			args{
				timeOut:    1,
				flagOutput: &[]string{"./testdata/report.json"}[0],
			},
		},
		{
			"Run Main - URL and Visibility Flag set",
			args{
				timeOut:        1,
				flagURL:        &[]string{testFileServer.URL + "/test.zip"}[0],
				flagVisibility: &[]string{"public"}[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Parse the flags.
			bParseFlags = tt.args.parseFlags

			// Alternate config.
			if tt.args.altConfig != nil {
				oldConf := procCfg
				procCfg = *tt.args.altConfig
				defer func() {
					procCfg = oldConf
				}()
			}

			// -output
			if tt.args.flagOutput != nil && *tt.args.flagOutput != "" {
				flagOutput = tt.args.flagOutput
			}

			// -url
			if tt.args.flagURL != nil && *tt.args.flagURL != "" {
				flagURL = tt.args.flagURL
			}

			// -visibility
			if tt.args.flagVisibility != nil && *tt.args.flagVisibility != "" {
				flagVisibility = tt.args.flagVisibility
			}

			if tt.args.version {
				bVersion = &[]bool{true}[0]
				Version = "0.0.1-test"
				Build = "12345"
			}

			if tt.args.authError {
				oldID := serviceConfig["tide"]["key"]
				serviceConfig["tide"]["key"] = "error"
				defer func() {
					serviceConfig["tide"]["key"] = oldID
				}()
			}

			// Run as goroutine and wait for terminate signal.
			go main()

			if tt.args.messageChannel != nil {
				oldCMessage := cMessage
				cMessage = tt.args.messageChannel
				cMessage <- tt.args.msg
				defer func() {
					cMessage = oldCMessage
				}()
			}

			// Sleep for one second. Allows for one poll action.
			time.Sleep(time.Millisecond * 100 * tt.args.timeOut)
			terminateChannel <- struct{}{}
		})
	}
}

func resetServiceConfig() {
	serviceConfig = getServiceConfig()
}

func Test_pollProvider(t *testing.T) {
	type args struct {
		c        chan message.Message
		provider message.Provider
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Poll Messages",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"message",
				},
			},
		},
		{
			"Poll Messages - Critical Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"critical",
				},
			},
		},
		{
			"Poll Messages - Quota Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"quota",
				},
			},
		},
		{
			"Poll Messages - Message Length Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"lenError",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pollProvider(tt.args.c, tt.args.provider)
		})
	}
}

func Test_processMessage(t *testing.T) {

	doProcess = MockDoProcess
	defer func() {
		doProcess = executeProcessFunc
	}()

	cMessageProvider := messageProvider
	messageProvider = &mockMessageProvider{}
	defer func() { messageProvider = cMessageProvider }()

	tests := []struct {
		name    string
		msg     message.Message
		wantErr bool
	}{
		{
			"Ingest - Success",
			message.Message{
				Slug: "ingestSuccess",
			},
			false,
		},
		{
			"Ingest - Fail",
			message.Message{
				Slug: "ingestFail",
			},
			true,
		},
		{
			"Info - Success",
			message.Message{
				Slug: "infoSuccess",
			},
			false,
		},
		{
			"Info - Fail",
			message.Message{
				Slug: "infoFail",
			},
			true,
		},
		{
			"Lighthouse - Success",
			message.Message{
				Slug: "lhSuccess",
			},
			false,
		},
		{
			"Lighthouse - Fail",
			message.Message{
				Slug: "lhFail",
			},
			true,
		},
		{
			"Response - Success",
			message.Message{
				Slug: "resSuccess",
			},
			false,
		},
		{
			"Response - Fail",
			message.Message{
				Slug: "resFail",
			},
			true,
		},
		{
			"Message Success - Provider Success",
			message.Message{
				Slug: "resultSuccess",
			},
			false,
		},
		{
			"Message Success - Provider Fail",
			message.Message{
				Slug:        "resultSuccess",
				ExternalRef: &[]string{"removeFail"}[0],
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(1)
			if err := processMessage(tt.msg, &wg); (err != nil) != tt.wantErr {
				t.Errorf("processMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getStorageProvider(t *testing.T) {
	type args struct {
		config map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			"No Provider",
			args{},
			nil,
		},
		{
			"S3 Provider",
			args{
				map[string]map[string]string{
					"app": {
						"storage_provider": "s3",
					},
					"aws": {
						"key":       "",
						"secret":    "",
						"s3_region": "",
						"s3_bucket": "",
					},
				},
			},
			reflect.TypeOf(&s3.Provider{}),
		},
		{
			"GCS Provider",
			args{
				map[string]map[string]string{
					"app": {
						"storage_provider": "gcs",
					},
					"gcp": {
						"project":    "",
						"gcs_bucket": "",
					},
				},
			},
			reflect.TypeOf(&gcs.Provider{}),
		},
		{
			"Local Provider",
			args{
				map[string]map[string]string{
					"app": {
						"storage_provider": "local",
						"server_path":      "./testdata",
						"local_path":       "subdir",
					},
				},
			},
			reflect.TypeOf(&local.Provider{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStorageProvider(tt.args.config); reflect.TypeOf(got) != tt.want {
				t.Errorf("getStorageProvider() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}

func Test_getMessageProvider(t *testing.T) {
	type args struct {
		config map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			"Nil Provider",
			args{
				make(map[string]map[string]string),
			},
			nil,
		},
		{
			"SQS Provider",
			args{
				map[string]map[string]string{
					"app": {
						"message_provider": "sqs",
					},
					"aws": {
						"key":        "",
						"secret":     "",
						"sqs_region": "",
						"sqs_queue":  "",
						"s3_region":  "",
					},
				},
			},
			reflect.TypeOf(&sqs.Provider{}),
		},
		{
			"Firestore Provider",
			args{
				map[string]map[string]string{
					"app": {
						"message_provider": "firestore",
					},
					"gcp": {
						"project":   "mock-project-id",
						"gcf_queue": "test-queue",
					},
				},
			},
			reflect.TypeOf(&firestore.Provider{}),
		},
		{
			"Mongo Provider",
			args{
				map[string]map[string]string{
					"app": {
						"message_provider": "local",
					},
					"mongo": {
						"user":     "test",
						"pass":     "test",
						"host":     "localhost:27017",
						"database": "test-db",
						"queue":    "test-queue",
					},
				},
			},
			reflect.TypeOf(&mongo.Provider{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessageProvider(tt.args.config); reflect.TypeOf(got) != tt.want {
				t.Errorf("getMessageProvider() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
