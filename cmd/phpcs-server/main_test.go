package main

import (
	"reflect"
	"testing"

	"github.com/wptide/pkg/message"
	tideApi "github.com/wptide/pkg/tide"
	"os"
	"errors"
	"github.com/wptide/pkg/env"
	"time"
	"log"
	"bytes"
	"github.com/wptide/pkg/audit"
	"net/http/httptest"
	"net/http"
)

var (
	envTest = map[string]string{
		// Tide API config
		"API_AUTH_URL":            "http://tide.local/api/tide/v1/auth",
		"API_HTTP_HOST":           "tide.local",
		"API_PROTOCOL":            "http",
		"API_KEY":                 "tideapikey",
		"API_SECRET":              "tideapisecret",
		"API_VERSION":             "v1",
		//
		// AWS API settings
		"AWS_API_KEY":             "sqskey",
		"AWS_API_SECRET":          "sqssecret",
		//
		// AWS S3 settings
		"AWS_S3_BUCKET_NAME":      "test-bucket",
		"AWS_S3_REGION":           "us-west-2",
		//
		// AWS SQS settings
		"AWS_SQS_QUEUE_LH":        "test-queue",
		"AWS_SQS_REGION":          "us-west-2",
		"AWS_SQS_VERSION":         "2012-11-05",
		//
		// PHPCS Server settings
		"PHPCS_CONCURRENT_AUDITS": "1",
	}
)

var fileServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
}))

type mockMessageProvider struct {
	Type string
}

func (m mockMessageProvider) SendMessage(msg *message.Message) error {
	return nil
}
func (m mockMessageProvider) GetNextMessage() (*message.Message, error) {
	switch m.Type {
	case "critical":
		pErr := message.NewProviderError("Critical issue.")
		pErr.Type = message.ErrCritcal
		return nil, pErr
	case "quota":
		pErr := message.NewProviderError("Over quota.")
		pErr.Type = message.ErrOverQuota
		return nil, pErr
	case "message":
		return &message.Message{}, nil
	case "lenError":
		return &message.Message{Title: "lenError"}, nil
	default:
		return nil, nil
	}
}
func (m mockMessageProvider) DeleteMessage(ref *string) error {
	return nil
}

type mockTide struct {
	apiError bool
}

func (m mockTide) Authenticate(clientId, clientSecret, authEndpoint string) error {
	if clientId == "error" {
		return errors.New("something went wrong")
	}
	return nil
}

func (m mockTide) SendPayload(method, endpoint, data string) (string, error) {

	if m.apiError {
		return "", errors.New("API error")
	}

	return "", nil
}

type mockProcess struct{}

func (m mockProcess) Process(msg message.Message, result *audit.Result) {
	r := *result;
	if msg.Title == "lenError" {
		r["mockError"] = errors.New("something went wrong")
	}
}
func (m mockProcess) Kind() string {
	return "mock"
}

func Test_main(t *testing.T) {

	tempWriter := bytes.Buffer{}
	log.SetOutput(&tempWriter)

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
		messageChannel chan *message.Message
		timeOut        time.Duration
		msg            *message.Message
		parseFlags     bool
		version        bool
		authError      bool
		flagUrl        *string
		flagOutput     *string
		flagVisibility *string
	}

	tests := []struct {
		name string
		args args
	}{
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
				messageChannel: make(chan *message.Message, 1),
				timeOut:        1,
				msg:            &message.Message{},
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
				flagUrl:        &[]string{fileServer.URL + "/test.zip"}[0],
				flagVisibility: &[]string{"public"}[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Parse the flags.
			bParseFlags = tt.args.parseFlags

			// -output
			if tt.args.flagOutput != nil && *tt.args.flagOutput != "" {
				flagOutput = tt.args.flagOutput
			}

			// -url
			if tt.args.flagUrl != nil && *tt.args.flagUrl != "" {
				flagUrl = tt.args.flagUrl
				cMessage = nil
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
				tideConfig.id = "error"
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
			time.Sleep(time.Millisecond * tt.args.timeOut)
			terminateChannel <- struct{}{}
		})
	}
}
func setupConfig() {
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
}

func Test_messageChannel(t *testing.T) {

	tempWriter := bytes.Buffer{}
	log.SetOutput(&tempWriter)

	type args struct {
		provider  message.MessageProvider
		buffer    chan struct{}
		processes []audit.Processor
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			"Successful Message",
			args{
				provider: &mockMessageProvider{
					Type: "message",
				},
			},
			reflect.TypeOf(make(chan *message.Message)),
		},
		{
			"Critical Error Message",
			args{
				provider: &mockMessageProvider{
					Type: "critical",
				},
			},
			reflect.TypeOf(make(chan *message.Message)),
		},
		{
			"Over Quota Message",
			args{
				provider: &mockMessageProvider{
					Type: "quota",
				},
			},
			reflect.TypeOf(make(chan *message.Message)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messageChannel(tt.args.provider, tt.args.buffer); !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("messageChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processMessage(t *testing.T) {

	type args struct {
		msg    *message.Message
		client tideApi.ClientInterface
		buffer <-chan struct{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"No Error Message",
			args{
				&message.Message{},
				&mockTide{},
				nil,
			},
		},
		{
			"Error Message",
			args{
				&message.Message{
					Title: "lenError",
				},
				&mockTide{},
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processes = []audit.Processor{
				&mockProcess{},
			}

			go processMessage(tt.args.msg, tt.args.client, tt.args.buffer)
			time.Sleep(time.Millisecond * 1)
		})
	}
}
