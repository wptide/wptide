package ingest

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

type Processor struct {}

func (p Processor) Process(msg message.Message, result *audit.Result) {
	r := *result
	r["ingest"] = "ingest"
}