package tide

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

type Processor struct {}

// Process takes *result and sends relevant results to the Tide API.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	r["tide"] = "tide"
}