package main

import (
	"context"
	"errors"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/message/mongo"
	"github.com/wptide/pkg/wporg"
)

type mongoDispatcher struct {
	ctx      context.Context
	user     string
	pass     string
	host     string
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
		if !ok {
			return errors.New("could not access queue: make sure the dispatcher is initialised")
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
		if !ok {
			messageProvider, _ = mongo.New(m.ctx, m.user, m.pass, m.host, m.database, collection.Collection, nil)
			m.providers[collectionID] = messageProvider
		}
	}
	return nil
}

func (m *mongoDispatcher) Close() error {
	for collectionID := range m.Collections {
		messageProvider, ok := m.providers[collectionID]
		if ok {
			messageProvider.Close()
		}
	}
	return nil
}
