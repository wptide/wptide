package main

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/wptide/pkg/env"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/payload"
)

var (
	envTest = map[string]string{
		// Tide API config
		"API_AUTH_URL":        "http://tide.local/api/tide/v1/auth",
		"API_HTTP_HOST":       "tide.local",
		"API_PROTOCOL":        "http",
		"API_KEY":             "tideapikey",
		"API_SECRET":          "tideapisecret",
		"API_VERSION":         "v1",
		//
		// AWS API settings
		"AWS_API_KEY":         "sqskey",
		"AWS_API_SECRET":      "sqssecret",
		//
		// AWS S3 settings
		"AWS_S3_BUCKET_NAME":   "test-bucket",
		"AWS_S3_REGION":        "us-west-2",
		//
		// AWS SQS settings
		"AWS_SQS_QUEUE_LH":     "test-queue",
		"AWS_SQS_REGION":       "us-west-2",
		"AWS_SQS_VERSION":      "2012-11-05",
		//
		// LH Server settings
		"LH_CONCURRENT_AUDITS": "1",
	}
)

func Test_initProcesses(t *testing.T) {

	type args struct {
		source <-chan message.Message
		config processConfig
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Valid Processes",
			args{
				make(chan message.Message),
				processConfig{
					igTempFolder:      "./testdata/tmp",
					lhTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			false,
		},
		{
			"No Config",
			args{
				source: make(chan message.Message),
			},
			true,
		},
		{
			"No Source",
			args{
				config: processConfig{
					igTempFolder:      "./testdata/tmp",
					lhTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			true,
		},
		{
			"Ingest temp folder missing",
			args{
				make(chan message.Message),
				processConfig{
					lhTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			true,
		},
		{
			"Lighthouse temp folder missing",
			args{
				make(chan message.Message),
				processConfig{
					igTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			true,
		},
		{
			"Lighthouse storage provider missing",
			args{
				make(chan message.Message),
				processConfig{
					igTempFolder: "./testdata/tmp",
					lhTempFolder: "./testdata/tmp",
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			true,
		},
		{
			"Response payloaders missing",
			args{
				make(chan message.Message),
				processConfig{
					igTempFolder:      "./testdata/tmp",
					lhTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
				},
			},
			true,
		},
		{
			"Valid Processes with messages",
			args{
				make(chan message.Message),
				processConfig{
					igTempFolder:      "./testdata/tmp",
					lhTempFolder:      "./testdata/tmp",
					lhStorageProvider: &mockStorage{},
					resPayloaders: map[string]payload.Payloader{
						"tide": &mockPayloader{},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := initProcesses(tt.args.source, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("initProcesses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

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

	setupConfig()

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
		flagUrl        *string
		flagOutput     *string
		flagVisibility *string
		altConfig      *processConfig
	}

	tests := []struct {
		name     string
		args     args
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
				flagUrl:        &[]string{testFileServer.URL + "/test.zip"}[0],
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
			if tt.args.flagUrl != nil && *tt.args.flagUrl != "" {
				flagUrl = tt.args.flagUrl
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
				oldId := tideConfig.id
				tideConfig.id = "error"
				defer func() {
					tideConfig.id = oldId
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

func setupConfig() {
	// Setup lhConfig
	lhConfig = struct {
		region string
		key    string
		secret string
		queue  string
	}{
		env.GetEnv("AWS_SQS_REGION", ""),
		env.GetEnv("AWS_API_KEY", ""),
		env.GetEnv("AWS_API_SECRET", ""),
		env.GetEnv("AWS_SQS_QUEUE_LH", ""),
	}

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

	lhS3Config = struct {
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
}

func Test_pollProvider(t *testing.T) {
	type args struct {
		c        chan message.Message
		provider message.MessageProvider
		buffer   chan struct{}
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
				make(chan struct{}, 1),
			},
		},
		{
			"Poll Messages - Critical Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"critical",
				},
				make(chan struct{}, 1),
			},
		},
		{
			"Poll Messages - Quota Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"quota",
				},
				make(chan struct{}, 1),
			},
		},
		{
			"Poll Messages - Message Length Error",
			args{
				make(chan message.Message),
				&mockMessageProvider{
					"lenError",
				},
				make(chan struct{}, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pollProvider(tt.args.c, tt.args.provider, tt.args.buffer)
		})
	}
}
