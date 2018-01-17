package info

import (
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
)

func TestProcessor_Process(t *testing.T) {
	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name string
		p    Processor
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)
		})
	}
}
