package main

import (
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/wporg"
)

func getDispatchMessage(project wporg.RepoProject) *message.Message {
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
		PayloadType:         "tide",
	}

	switch project.Type {
	case "themes":
		msg.SourceURL = "https://downloads.wordpress.org/theme/" + project.Slug + "." + project.Version + ".zip"
		msg.ProjectType = "theme"
		msg.Content = project.Description
		msg.Standards = append(msg.Standards, "lighthouse")
		msg.Audits = append(msg.Audits, &message.Audit{
			Type:    "lighthouse",
			Options: &message.AuditOption{},
		})
	case "plugins":
		msg.SourceURL = project.DownloadLink
		msg.ProjectType = "plugin"
		msg.Content = project.ShortDescription
	}
	return msg
}
