package main

import (
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/wporg"
	"github.com/wptide/pkg/message/sqs"
	"errors"
)

type sqsDispatcher struct {
	Queues map[string]struct {
		Region   string
		Key      string
		Secret   string
		Queue    string
		Endpoint string
		Active   bool
		Accepts  string // "all" or "themes" or "plugins"
	}
	providers map[string]message.MessageProvider
}

func (s sqsDispatcher) Dispatch(project wporg.RepoProject) error {

	msg := &message.Message{
		ResponseAPIEndpoint: apiEndpoint,
		Title:               project.Name,
		Slug:                project.Slug,
		SourceType:          "zip",
		RequestClient:       apiClient,
		Force:               forcedSync,
		Visibility:          auditVisibility,
		Standards:           defaultStandards(),
		Audits:              defaultAudits(),
	}

	switch project.Type {
	case "themes":
		msg.SourceURL = "https://downloads.wordpress.org/theme/" + project.Slug + "." + project.Version + ".zip"
		msg.ProjectType = "theme"
		msg.Content = project.Description
	case "plugins":
		msg.SourceURL = project.DownloadLink
		msg.ProjectType = "plugin"
		msg.Content = project.ShortDescription
	}

	for queueID, queue := range s.Queues {

		queueProvider, ok := s.providers[queueID]
		if ! ok {
			return errors.New("Could not access queue. Make sure the dispatcher is initialised.")
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
		if ! ok {
			queueProvider = sqs.NewSqsProvider(queue.Region, queue.Key, queue.Secret, queue.Queue)
			s.providers[queueID] = queueProvider
		}
	}
	return nil
}
