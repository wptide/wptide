package main

import (
	"os"
	"testing"
	"time"
	commonProcess "github.com/xwp/go-tide/src/process"
	"github.com/wptide/pkg/env"
	"github.com/wptide/pkg/message"
	"bytes"
	"log"
	"sync"
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
		"AWS_SQS_QUEUE_PHPCS": "test-queue",
		"AWS_SQS_REGION":      "us-west-2",
		"AWS_SQS_VERSION":     "2012-11-05",
		//
		// PHPCS Server settings
		"PHPCS_CONCURRENT_AUDITS": "1",
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
		flagWPRuleset  *string
		altConfig      *processConfig
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"Run Main - Auth Error",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				authError:      true,
			},
		},
		{
			"Run Main",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				version:        true,
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
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				parseFlags:     true,
			},
		},
		{
			"Run Main - Output Flag set",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				flagOutput:     &[]string{"./testdata/report.json"}[0],
			},
		},
		{
			"Run Main - URL and Visibility Flag set",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				flagUrl:        &[]string{testFileServer.URL + "/test.zip"}[0],
				flagVisibility: &[]string{"public"}[0],
			},
		},
		{
			"Run Main - WP PHPCompatibility Flag set",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        1,
				flagWPRuleset:  &[]string{"./testdata/wordpress.xml"}[0],
			},
		},
		{
			"Run Main - Process Error",
			args{
				messageChannel: make(chan message.Message, 1),
				timeOut:        0,
				version:        true,
				// Invalid config. Will cause a process.Run() error.
				altConfig: &processConfig{
					igTempFolder:         "./testdata/tmp",
					phpcsTempFolder:      "./testdata/tmp",
					phpcsStorageProvider: &mockStorage{},
				},
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
				procCfg.igTempFolder = tt.args.altConfig.igTempFolder
				procCfg.phpcsTempFolder = tt.args.altConfig.phpcsTempFolder
				procCfg.phpcsStorageProvider = tt.args.altConfig.phpcsStorageProvider

				if tt.args.altConfig.igTempFolder == "no-proc" {
					procCfg = nil
				}

				defer func() {
					procCfg = &processConfig{}
					procCfg.igTempFolder = oldConf.igTempFolder
					procCfg.phpcsTempFolder = oldConf.phpcsTempFolder
					procCfg.phpcsStorageProvider = oldConf.phpcsStorageProvider
					procCfg.resPayloaders = oldConf.resPayloaders
				}()
			}

			// -output
			if tt.args.flagOutput != nil && *tt.args.flagOutput != "" {
				flagOutput = tt.args.flagOutput
			}

			// -url
			if tt.args.flagUrl != nil && *tt.args.flagUrl != "" {
				flagUrl = tt.args.flagUrl
				defer func() {
					flagUrl = &[]string{""}[0]
				}()
			}

			// -visibility
			if tt.args.flagVisibility != nil && *tt.args.flagVisibility != "" {
				flagVisibility = tt.args.flagVisibility
				defer func() {
					flagVisibility = &[]string{""}[0]
				}()
			}

			// If -wp-phpcompat-ruleset was provided then override the ruleset path.
			if tt.args.flagWPRuleset != nil {
				oldFlag := flagWPRuleset
				flagWPRuleset = tt.args.flagWPRuleset
				defer func() {
					flagWPRuleset = oldFlag
				}()
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

			// Sleep before terminating.
			time.Sleep(time.Millisecond * 300 * tt.args.timeOut)
			terminateChannel <- struct{}{}
		})
	}
}

func setupConfig() {
	// Setup queueConfig
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
}

func Test_pollProvider(t *testing.T) {
	type args struct {
		c        chan message.Message
		provider message.MessageProvider
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

	doIngest = MockDoIngest
	doInfo = MockDoInfo
	doPhpcs = MockDoPhpcs
	doResponse = MockDoResponse
	defer func() {
		doIngest = commonProcess.DoIngest
		doInfo = commonProcess.DoInfo
		doPhpcs = commonProcess.DoPhpcs
		doResponse = commonProcess.DoResponse
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
			"Phpcs - Success",
			message.Message{
				Slug: "phpcsSuccess",
			},
			false,
		},
		{
			"Phpcs - Fail",
			message.Message{
				Slug: "phpcsFail",
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
