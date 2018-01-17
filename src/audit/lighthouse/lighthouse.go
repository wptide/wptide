package lighthouse

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

type Processor struct {}

// Process will run a lighthouse audit against a theme and to *results.
// Note: It checks *result first to see if the message has been ingested and if its processing a theme.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	r["lighthouse"] = "lighthouse"
}