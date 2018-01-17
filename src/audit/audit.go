package audit

import "github.com/xwp/go-tide/src/message"

// Result is an interface map used to store results from processes.
type Result map[string]interface{}

// Processor is an interface for processing a message and producing results.
type Processor interface{
	Process(msg message.Message, result *Result)
}