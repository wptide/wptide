package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nanobox-io/golang-scribble"
	"github.com/wptide/pkg/wporg"
)

var mutex = sync.Mutex{}

type scribbleChecker struct {
	db *scribble.Driver
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

	err := s.db.Write(project.Type, project.Slug, project)
	return err
}

func (s *scribbleChecker) SetSyncTime(event, projectType string, t time.Time) {
	key := fmt.Sprintf("%s-sync-%s", projectType, event)
	s.db.Write("info", key, t.UnixNano())
}

func (s scribbleChecker) GetSyncTime(event, projectType string) time.Time {
	key := fmt.Sprintf("%s-sync-%s", projectType, event)
	var timestamp int64
	var t time.Time

	err := s.db.Read("info", key, &timestamp)
	if err != nil {
		t, _ = time.Parse(wporg.TimeFormat, "1970-01-01 12:00am MST")
	} else {
		t = time.Unix(0, timestamp)
	}

	return t
}

func newScribbleChecker(path string) *scribble.Driver {
	db, _ := scribble.New(path, nil)
	return db
}
