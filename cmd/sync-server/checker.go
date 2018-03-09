package main

import (
	"github.com/wptide/pkg/wporg"
	"github.com/nanobox-io/golang-scribble"
	"errors"
	"sync"
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

func (s scribbleChecker) RecordUpdate(project wporg.RepoProject) error {
	mutex.Lock()
	defer mutex.Unlock()

	if s.db == nil {
		return errors.New("no db to write to")
	}

	err := s.db.Write(project.Type, project.Slug, project)
	return err
}

func newScribbleChecker(path string) *scribble.Driver {
	db, _ := scribble.New(path, nil)
	return db
}