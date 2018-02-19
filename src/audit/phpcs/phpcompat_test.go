package phpcs

import (
	"io"
	"testing"

	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"os"
	"github.com/wptide/pkg/tide"
)

func TestPhpCompat_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := PhpCompat{}
		if got := p.Kind(); got != "phpcs_phpcompat" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be phpcs_phpcompat.", got)
		}
	})
}

func TestPhpCompat_Process(t *testing.T) {

	parent := &Processor{
		Standard: "phpcompatibility",
	}
	readerCompatSet1, _ := os.Open("./testdata/samples/phpcompat1.json")
	readerCompatSet2, _ := os.Open("./testdata/samples/phpcompat2.json")
	//readerFailed, _ := os.Open(testFiles["fail"])
	//readerInvalid, _ := os.Open(testFiles["invalid"])

	storageProvider = &mockStorage{}

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
			"PhpCompatibility - Set 1",
			fields{
				Report:        readerCompatSet1,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":  &storageProvider,
					"tempFolder": "/tmp",
					"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Set 2",
			fields{
				Report:        readerCompatSet2,
				ParentProcess: parent,
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":  &storageProvider,
					"tempFolder": "/tmp",
					"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PhpCompat{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.Process(tt.args.msg, tt.args.result)
		})
	}
}

func TestPhpCompat_SetReport(t *testing.T) {

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
			p := &PhpCompat{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.SetReport(tt.args.report)
		})
	}
}

func TestPhpCompat_Parent(t *testing.T) {
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
			p := &PhpCompat{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
			}
			p.Parent(tt.args.parent)
		})
	}
}
