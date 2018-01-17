package main

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/audit/ingest"
	"github.com/xwp/go-tide/src/audit/info"
	"github.com/xwp/go-tide/src/audit/lighthouse"
	"github.com/xwp/go-tide/src/audit/tide"
)

func main() {
}

// processMessage takes an SQS message and runs it through an audit process.
func processMessage(msg message.Message) {

	// An slice of processes that need to be performed on the message.
	// A slice ensures that they happen in the correct order.
	processes := []audit.Processor{
		&ingest.Processor{},
		&info.Processor{},
		&lighthouse.Processor{},
		&tide.Processor{},
	}

	// Initialise an empty result.
	result := &audit.Result{}

	// Run through each processor and update the result.
	// Note: The result is a pointer so we're passing by reference.
	for _, proc := range processes {
		proc.Process(msg, result)
	}
}
