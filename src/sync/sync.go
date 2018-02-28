package sync

import (
	"github.com/wptide/pkg/message"
	"github.com/xwp/go-tide/src/wporg"
)

type Synchronizer struct {

}

type Dispatcher interface{
	Dispatch(project wporg.RepoProject) error
	Init() error
}

type UpdateChecker interface {
	UpdateCheck(project wporg.RepoProject) bool
	RecordUpdate(project wporg.RepoProject) error
}

type MessagesProcessor interface {
	ProcessMessages(messages chan *message.Message) error
}