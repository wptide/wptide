package ingest

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/util"
)

type Processor struct{}

func (p Processor) Kind() string {
	return "ingest"
}

// Process takes *result and sends relevant results to the Tide API.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// Only ingest if we do not know if this is a theme or we don't have a specific
	// route given. This will have to be calculated from checksum if not.
	isCollection := util.IsCollectionEndpoint(msg.ResponseAPIEndpoint)

	if ! isCollection {
		return
	}



	r["ingest"] = "ingest"
}