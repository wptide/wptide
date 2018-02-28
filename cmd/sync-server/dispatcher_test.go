package main

import (
	"testing"

	"github.com/wptide/pkg/message"
	"github.com/xwp/go-tide/src/wporg"
	"errors"
	"log"
	"bytes"
)

var (
	testQueues = map[string]struct {
		Region   string
		Key      string
		Secret   string
		Queue    string
		Endpoint string
		Active   bool
		Accepts  string
	}{
		"queue1": {
			"region1",
			"thekey",
			"thesecret",
			"primary",
			"http://test.local",
			true,
			"all",
		},
		"themes": {
			"region1",
			"thekey",
			"thesecret",
			"themes",
			"http://test.local",
			true,
			"themes",
		},
		"plugins": {
			"region1",
			"thekey",
			"thesecret",
			"plugins",
			"http://test.local",
			true,
			"plugins",
		},
		"plugins_disables": {
			"region1",
			"thekey",
			"thesecret",
			"plugins_disabled",
			"http://test.local",
			false,
			"plugins",
		},
	}
)

// Setup a mockProvider.
type mockProvider struct{}

func (m mockProvider) SendMessage(msg *message.Message) error {
	if msg.Slug == "error" {
		return errors.New("something went wrong")
	}
	return nil
}
func (m mockProvider) GetNextMessage() (*message.Message, error) {
	return nil, nil
}
func (m mockProvider) DeleteMessage(ref *string) error {
	return nil
}

func initMockProviders(s *sqsDispatcher) {
	for queueID, _ := range s.Queues {
		queueProvider, ok := s.providers[queueID]
		if ! ok {
			queueProvider = &mockProvider{}
			s.providers[queueID] = queueProvider
		}
	}
}

func Test_sqsDispatcher_Dispatch(t *testing.T) {

	// Don't log to os.Stdin.
	b := bytes.Buffer{}
	log.SetOutput(&b)

	type fields struct {
		Queues map[string]struct {
			Region   string
			Key      string
			Secret   string
			Queue    string
			Endpoint string
			Active   bool
			Accepts  string
		}
		providers map[string]message.MessageProvider
	}
	type args struct {
		project wporg.RepoProject
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		skipInit bool
		wantErr  bool
	}{
		{
			"Theme Project",
			fields{
				testQueues,
				make(map[string]message.MessageProvider),
			},
			args{
				wporg.RepoProject{
					Type: "themes",
				},
			},
			false,
			false,
		},
		{
			"Plugin Project",
			fields{
				testQueues,
				make(map[string]message.MessageProvider),
			},
			args{
				wporg.RepoProject{
					Type: "plugins",
				},
			},
			false,
			false,
		},
		{
			"Theme Project - No Init()",
			fields{
				testQueues,
				make(map[string]message.MessageProvider),
			},
			args{
				wporg.RepoProject{
					Type: "themes",
				},
			},
			true,
			true,
		},
		{
			"Theme Project - Provider Fail",
			fields{
				testQueues,
				make(map[string]message.MessageProvider),
			},
			args{
				wporg.RepoProject{
					Type: "themes",
					Slug: "error",
				},
			},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqsDispatcher{
				Queues:    tt.fields.Queues,
				providers: tt.fields.providers,
			}

			if ! tt.skipInit {
				initMockProviders(s)
			}

			if err := s.Dispatch(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("sqsDispatcher.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqsDispatcher_Init(t *testing.T) {
	type fields struct {
		Queues map[string]struct {
			Region   string
			Key      string
			Secret   string
			Queue    string
			Endpoint string
			Active   bool
			Accepts  string
		}
		providers map[string]message.MessageProvider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Init Providers",
			fields{
				testQueues,
				make(map[string]message.MessageProvider),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqsDispatcher{
				Queues:    tt.fields.Queues,
				providers: tt.fields.providers,
			}
			if err := s.Init(); (err != nil) != tt.wantErr {
				t.Errorf("sqsDispatcher.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
