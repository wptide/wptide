package audit

import (
	"testing"
	"github.com/xwp/go-tide/src/message"
)

type mockProcess struct{}

func (m mockProcess) Process(msg message.Message, result *Result) {
	return
}

func (m mockProcess) Kind() string {
	return "mock"
}

func TestCanRunAudit(t *testing.T) {
	type args struct {
		p       Processor
		results *Result
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Can Run Process",
			args{
				mockProcess{},
				&Result{
					"audits": []string{"mock"},
				},
			},
			true,
		},
		{
			"Cannot Run Process",
			args{
				mockProcess{},
				&Result{
					"audits": []string{"other"},
				},
			},
			false,
		},
		{
			"No audits",
			args{
				mockProcess{},
				&Result{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanRunAudit(tt.args.p, tt.args.results); got != tt.want {
				t.Errorf("CanRunAudit() = %v, want %v", got, tt.want)
			}
		})
	}
}
