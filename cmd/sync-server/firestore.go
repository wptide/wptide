package main

import (
	"context"
	"errors"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/message/firestore"
	"github.com/wptide/pkg/wporg"
)

type firestoreDispatcher struct {
	ProjectID   string
	Collections map[string]struct {
		Collection string
		Active     bool
		Accepts    string // "all" or "themes" or "plugins".
	}
	providers map[string]message.Provider
}

func (fs firestoreDispatcher) Dispatch(project wporg.RepoProject) error {

	msg := getDispatchMessage(project)

	for collectionID, collection := range fs.Collections {
		queueProvider, ok := fs.providers[collectionID]
		if !ok {
			return errors.New("could not access queue: make sure the dispatcher is initialized")
		}
		if collection.Active && (collection.Accepts == project.Type || collection.Accepts == "all") {
			err := queueProvider.SendMessage(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (fs *firestoreDispatcher) Init() error {
	for collectionID, collection := range fs.Collections {
		messageProvider, ok := fs.providers[collectionID]
		if !ok {
			messageProvider, _ = firestore.New(context.Background(), fs.ProjectID, collection.Collection)
			fs.providers[collectionID] = messageProvider
		}
	}
	return nil
}

func (fs *firestoreDispatcher) Close() error {
	for collectionID := range fs.Collections {
		if messageProvider, ok := fs.providers[collectionID]; ok {
			messageProvider.Close()
		}
	}
	return nil
}
