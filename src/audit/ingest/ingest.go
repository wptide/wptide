package ingest

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

type Processor struct {}

// Process takes the SQS message and downloads the requested plugin/theme for further processing.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	r["ingest"] = "ingest"
}