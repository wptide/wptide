package main

import (
	"github.com/wptide/pkg/wporg"
	"github.com/nanobox-io/golang-scribble"
	"errors"
	"sync"
	"time"
)

var mutex = sync.Mutex{}

type scribbleChecker struct {
	db          *scribble.Driver
	lastUpdated *wporg.RepoProject
}

func (s scribbleChecker) UpdateCheck(project wporg.RepoProject) bool {
	record := wporg.RepoProject{}

	if s.db == nil {
		return false
	}

	// Attempt to get current entry.
	s.db.Read(project.Type, project.Slug, &record)
	return record.LastUpdated != project.LastUpdated || record.Version != project.Version
}

func (s *scribbleChecker) RecordUpdate(project wporg.RepoProject) error {
	mutex.Lock()
	defer mutex.Unlock()

	if s.db == nil {
		return errors.New("no db to write to")
	}

	lastDate, _ := time.Parse(wporg.TimeFormat, s.lastUpdated.LastUpdated)
	thisDate, _ := time.Parse(wporg.TimeFormat, project.LastUpdated)

	if thisDate.After(lastDate) {
		s.setLastUpdated(project)
		s.db.Write("info", "last-updated", project)
	}

	err := s.db.Write(project.Type, project.Slug, project)
	return err
}

func (s *scribbleChecker) GetLastUpdated() (*wporg.RepoProject, error) {
	record := wporg.RepoProject{}

	if s.db == nil {
		return nil, errors.New("no db to get last record from")
	}

	err := s.db.Read("info", "last-updated", &record)
	if err != nil {
		record = wporg.RepoProject{
			LastUpdated: "1970-01-01 12:00am MST",
		}
	}

	s.setLastUpdated(record)

	return &record, err
}

func (s *scribbleChecker) setLastUpdated(project wporg.RepoProject) {
	s.lastUpdated = &project
}

func newScribbleChecker(path string) *scribble.Driver {
	db, _ := scribble.New(path, nil)
	return db
}
