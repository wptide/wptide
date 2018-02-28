package sync

import (
	"github.com/xwp/go-tide/src/wporg"
)

// Dispatcher describes an interface for dispatching RepoProjects to
// a queue service. The service will need to implement this interface.
type Dispatcher interface {
	Dispatch(project wporg.RepoProject) error
	Init() error
}

// UpdateChecker describes an interface to determine and record the currency
// of the last dispatched RepoProject.
type UpdateChecker interface {
	UpdateCheck(project wporg.RepoProject) bool
	RecordUpdate(project wporg.RepoProject) error
}
