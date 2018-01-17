package tide

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
)

type Processor struct{}

// Process takes *result and sends relevant results to the Tide API.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// @TODO: Add environment vars for Tide API
	if r["client"] != nil {
		client := r["client"].(*tide.ClientInterface)
		p.sendResults(client,"","")
	}

	r["tide"] = "tide"
}

func (p Processor) sendResults(client *tide.ClientInterface, endpoint, data string) {
	c := *client
	c.SendPayload("POST", endpoint, data)
}
