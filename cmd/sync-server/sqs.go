package main

import (
	"errors"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/message/sqs"
	"github.com/wptide/pkg/wporg"
)

type sqsDispatcher struct {
	Queues map[string]struct {
		Region   string
		Key      string
		Secret   string
		Queue    string
		Endpoint string
		Active   bool
		Accepts  string // "all" or "themes" or "plugins".
	}
	providers map[string]message.Provider
}

func (s sqsDispatcher) Dispatch(project wporg.RepoProject) error {

	msg := getDispatchMessage(project)

	for queueID, queue := range s.Queues {

		queueProvider, ok := s.providers[queueID]
		if !ok {
			return errors.New("could not access queue: make sure the dispatcher is initialised")
		}

		if queue.Active && (queue.Accepts == project.Type || queue.Accepts == "all") {
			err := queueProvider.SendMessage(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *sqsDispatcher) Init() error {
	for queueID, queue := range s.Queues {
		queueProvider, ok := s.providers[queueID]
		if !ok {
			queueProvider = sqs.NewSqsProvider(queue.Region, queue.Key, queue.Secret, queue.Queue)
			s.providers[queueID] = queueProvider
		}
	}
	return nil
}

func (s *sqsDispatcher) Close() error {
	for queueID := range s.Queues {
		queueProvider, ok := s.providers[queueID]
		if ok {
			queueProvider.Close()
		}
	}
	return nil
}
