package tide

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

func TestProcessor_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := Processor{}
		if got := p.Kind(); got != "tide" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be tide.", got)
		}
	})
}