package info

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

type Processor struct {}

// Process attempts to retrieve information about the plugin/theme.
// Note: It checks *result first to see if the message has been ingested.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	r["info"] = "info"
}