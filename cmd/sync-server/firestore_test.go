package main

import (
	"testing"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/wporg"
)

var (
	testCollections = map[string]struct {
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

func initMockFSProviders(fs *firestoreDispatcher) {
	for collectionID := range fs.Collections {
		queueProvider, ok := fs.providers[collectionID]
		if !ok {
			queueProvider = &mockProvider{}
			fs.providers[collectionID] = queueProvider
		}
	}
}

func Test_firestoreDispatcher_Dispatch(t *testing.T) {
	type fields struct {
		ProjectID   string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
		}
		providers map[string]message.Provider
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
				"fake-id",
				testCollections,
				make(map[string]message.Provider),
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
				"fake-id",
				testCollections,
				make(map[string]message.Provider),
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
				"fake-id",
				testCollections,
				make(map[string]message.Provider),
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
				"fake-id",
				testCollections,
				make(map[string]message.Provider),
			},
			args{
				wporg.RepoProject{
					Type: "themes",
					Slug: "error",
				},
			},
			false,
			true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &firestoreDispatcher{
				ProjectID:   tt.fields.ProjectID,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}

			if !tt.skipInit {
				initMockFSProviders(fs)
			}

			if err := fs.Dispatch(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("firestoreDispatcher.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_firestoreDispatcher_Init(t *testing.T) {
	type fields struct {
		ProjectID   string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
		}
		providers map[string]message.Provider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Init Providers",
			fields{
				"fake-id",
				testCollections,
				make(map[string]message.Provider),
			},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &firestoreDispatcher{
				ProjectID:   tt.fields.ProjectID,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}
			if err := fs.Init(); (err != nil) != tt.wantErr {
				t.Errorf("firestoreDispatcher.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_firestoreDispatcher_Close(t *testing.T) {
	type fields struct {
		ProjectID   string
		Collections map[string]struct {
			Collection string
			Active     bool
			Accepts    string
		}
		providers map[string]message.Provider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Close with collections",
			fields{
				"abc",
				testCollections,
				make(map[string]message.Provider),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &firestoreDispatcher{
				ProjectID:   tt.fields.ProjectID,
				Collections: tt.fields.Collections,
				providers:   tt.fields.providers,
			}

			initMockFSProviders(fs)

			if err := fs.Close(); (err != nil) != tt.wantErr {
				t.Errorf("firestoreDispatcher.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
