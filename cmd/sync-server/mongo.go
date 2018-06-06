package main

import (
	"github.com/wptide/pkg/wporg"
	"github.com/wptide/pkg/message"
	"context"
	"errors"
	"github.com/wptide/pkg/message/mongo"
)

type mongoDispatcher struct {
	ctx context.Context
	user string
	pass string
	host string
	database string

	Collections map[string]struct {
		Collection string
		Active     bool
		Accepts    string // "all" or "themes" or "plugins"
	}
	providers map[string]message.MessageProvider
}

func (m mongoDispatcher) Dispatch(project wporg.RepoProject) error {

	msg := getDispatchMessage(project)

	for collectionID, collection := range m.Collections {
		queueProvider, ok := m.providers[collectionID]
		if ! ok {
			return errors.New("Could not access queue. Make sure the dispatcher is initialised.")
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

func (m *mongoDispatcher) Init() error {
	for collectionID, collection := range m.Collections {
		messageProvider, ok := m.providers[collectionID]
		if ! ok {
			messageProvider, _ = mongo.New(m.ctx, m.user, m.pass, m.host, m.database, collection.Collection, nil)
			m.providers[collectionID] = messageProvider
		}
	}
	return nil
}

func (m *mongoDispatcher) Close() error {
	for collectionID, _ := range m.Collections {
		messageProvider, ok := m.providers[collectionID]
		if ok {
			messageProvider.Close()
		}
	}
	return nil
}