package main

import (
	"context"
	"testing"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/wporg"
)

var (
	mongoCollections = map[string]struct {
		Collection string
		Active     bool
		Accepts    string // "all" or "themes" or "plugins"
	}{
		"queue1": {
			"region1",
			true,
			"all",
		},
		"themes": {
			"region1",
			true,
			"themes",
		},
		"plugins": {
			"region1",
			true,
			"plugins",
		},
		"plugins_disables": {
			"region1",
			false,
			"plugins",
		},
	}
)

func initMockMongoProviders(m mongoDispatcher) {
	for collectionID, _ := range m.Collections {
		queueProvider, ok := m.providers[collectionID]
		if !ok {
			queueProvider = &mockProvider{}
			m.providers[collectionID] = queueProvider
		}
	}
}

func Test_mongoDispatcher_Dispatch(t *testing.T) {
	type fields struct {
		ctx         context.Context
		user        string
		pass        string
		host        string
		database    string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
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
				context.Background(),
				"",
				"",
				"mock-host:1234",
				"test-queue",
				mongoCollections,
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
				context.Background(),
				"",
				"",
				"mock-host:1234",
				"test-queue",
				mongoCollections,
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
				context.Background(),
				"",
				"",
				"mock-host:1234",
				"test-queue",
				mongoCollections,
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
				context.Background(),
				"",
				"",
				"mock-host:1234",
				"test-queue",
				mongoCollections,
				make(map[string]message.MessageProvider),
			},
			args{
				wporg.RepoProject{
					Type: "themes",
					Slug: "error",
				},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mongoDispatcher{
				ctx:         tt.fields.ctx,
				user:        tt.fields.user,
				pass:        tt.fields.pass,
				host:        tt.fields.host,
				database:    tt.fields.database,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}

			if !tt.skipInit {
				initMockMongoProviders(m)
			}

			if err := m.Dispatch(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("mongoDispatcher.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoDispatcher_Init(t *testing.T) {
	type fields struct {
		ctx         context.Context
		user        string
		pass        string
		host        string
		database    string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
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
				context.Background(),
				"",
				"",
				"mock-host:1234",
				"test-queue",
				mongoCollections,
				make(map[string]message.MessageProvider),
			},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mongoDispatcher{
				ctx:         tt.fields.ctx,
				user:        tt.fields.user,
				pass:        tt.fields.pass,
				host:        tt.fields.host,
				database:    tt.fields.database,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}
			if err := m.Init(); (err != nil) != tt.wantErr {
				t.Errorf("mongoDispatcher.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoDispatcher_Close(t *testing.T) {
	type fields struct {
		ctx         context.Context
		user        string
		pass        string
		host        string
		database    string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
		}
		providers map[string]message.MessageProvider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Close with collections",
			fields{
				context.Background(),
				"",
				"",
				"localhost:27017",
				"mock-db",
				mongoCollections,
				make(map[string]message.MessageProvider),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mongoDispatcher{
				ctx:         tt.fields.ctx,
				user:        tt.fields.user,
				pass:        tt.fields.pass,
				host:        tt.fields.host,
				database:    tt.fields.database,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}

			initMockMongoProviders(*m)

			if err := m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("mongoDispatcher.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
