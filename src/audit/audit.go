package audit

import "github.com/xwp/go-tide/src/message"

type Result map[string]interface{}

type Processor interface{
	Process(msg message.Message, result *Result)
}