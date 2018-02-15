package phpcs

import (
	"io"
	"testing"

	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"os"
	"github.com/wptide/pkg/tide"
)

var (
	testFiles = map[string]string{
		"cf7": "./testdata/samples/contactform7-75-phpcs_phpcompatibility.json",
		"t17": "./testdata/samples/twentyseventeen-14-phpcs_wordpress.json",
		"fail": "./testdata/nothing/here",
		"invalid": "./testdata/samples/invalid-json.json",
	}
)

func TestPhpcsSummary_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := PhpcsSummary{}
		if got := p.Kind(); got != "phpcs_summary" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be phpcs_summary.", got)
		}
	})
}

func TestPhpcsSummary_Process(t *testing.T) {

	parent := &Processor{
		Standard: "wordpress",
	}
	readerWordPress, _ := os.Open(testFiles["t17"])
	readerFailed, _ := os.Open(testFiles["fail"])
	readerInvalid, _ := os.Open(testFiles["invalid"])

	type fields struct {
		Report        io.Reader
		ParentProcess audit.Processor
	}
	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"PHPCS WordPress Standard",
			fields{
				Report:        readerWordPress,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "WordPress Standard",
				},
				&audit.Result{
					"phpcs_wordpress": &tide.AuditResult{},
				},
			},
		},
		{
			"PHPCS WordPress - No Parent Result",
			fields{
				Report:        readerWordPress,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "WordPress Standard",
				},
				&audit.Result{},
			},
		},
		{
			"PHPCS WordPress - No Report Reader/File",
			fields{
				Report:        readerFailed,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "WordPress Standard",
				},
				&audit.Result{},
			},
		},
		{
			"PHPCS WordPress - Invalid Report File",
			fields{
				Report:        readerInvalid,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "WordPress Standard",
				},
				&audit.Result{
					"phpcs_wordpress": &tide.AuditResult{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PhpcsSummary{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.Process(tt.args.msg, tt.args.result)
		})
	}
}

func TestPhpcsSummary_SetReport(t *testing.T) {

	parent := &Processor{}
	reader, _ := os.Open(testFiles["t17"])

	type fields struct {
		Report        io.Reader
		ParentProcess audit.Processor
	}
	type args struct {
		report io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Set Report Test",
			fields{
				ParentProcess: parent,
			},
			args{
				report: reader,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PhpcsSummary{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.SetReport(tt.args.report)
		})
	}
}

func TestPhpcsSummary_Parent(t *testing.T) {

	parent := &Processor{}
	reader, _ := os.Open(testFiles["t17"])

	type fields struct {
		Report        io.Reader
		ParentProcess audit.Processor
	}
	type args struct {
		parent audit.Processor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Set Parent Test",
			fields{
				Report: reader,
			},
			args{
				parent: parent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PhpcsSummary{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.Parent(tt.args.parent)
		})
	}
}
