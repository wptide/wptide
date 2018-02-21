package phpcs

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
)

func mockWriteFile(filename string, data []byte, perm os.FileMode) error {
	switch filename {
	case "/tmp/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff-phpcs_phpcompat-detail.json":
		return errors.New("something went wrong")
	default:
		return ioutil.WriteFile(filename, data, perm)
	}
}

func TestPhpCompat_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := PhpCompat{}
		if got := p.Kind(); got != "phpcs_phpcompat" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be phpcs_phpcompat.", got)
		}
	})
}

func TestPhpCompat_Process(t *testing.T) {

	// Proxy writeFile
	writeFile = mockWriteFile
	defer func() {
		writeFile = ioutil.WriteFile
	}()

	readerCompatSet1, _ := os.Open("./testdata/samples/phpcompat1.json")
	readerCompatSet2, _ := os.Open("./testdata/samples/phpcompat2.json")
	readerFailed, _ := os.Open(testFiles["fail"])
	readerInvalid, _ := os.Open(testFiles["invalid"])

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
				Report: readerCompatSet1,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Set 2",
			fields{
				Report: readerCompatSet2,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Cant write to disk",
			fields{
				Report: readerCompatSet1,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				},
			},
		},
		{
			"PhpCompatibility - Failed",
			fields{
				Report: readerFailed,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Invalid",
			fields{
				Report: readerInvalid,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Invalid Results from Parent",
			fields{
				Report: readerCompatSet1,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": "invalid results",
					"fileStore":              &storageProvider,
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
		{
			"PhpCompatibility - Upload fail",
			fields{
				Report: readerCompatSet1,
				ParentProcess: &Processor{
					Standard: "phpcompatibility",
				},
			},
			args{
				message.Message{
					Title: "PhpCompatibility Standard",
				},
				&audit.Result{
					"phpcs_phpcompatibility": &tide.AuditResult{},
					"fileStore":              "not a valid file store",
					"tempFolder":             "/tmp",
					"checksum":               "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Because we're reusing files...
			report, ok := tt.fields.Report.(io.ReadSeeker)
			if ok {
				// ... we have to "rewind" them.
				report.Seek(0, 0)
			}

			p := &PhpCompat{
				Report:        report,
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

func TestPhpCompat_summaryPath(t *testing.T) {
	type fields struct {
		Report        io.Reader
		ParentProcess audit.Processor
		resultPath    string
		resultFile    string
	}
	type args struct {
		r          audit.Result
		fileSuffix string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFilename string
		wantPath     string
		wantErr      bool
	}{
		{
			"With valid path",
			fields{
				Report: nil,
			},
			args{
				r: audit.Result{
					"tempFolder": "/tmp",
					"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
			"39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_phpcompat.json",
			"/tmp/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_phpcompat.json",
			false,
		},
		{
			"No tempFolder",
			fields{
				Report: nil,
			},
			args{
				r: audit.Result{
					"checksum": "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
			"",
			"",
			true,
		},
		{
			"No checksum",
			fields{
				Report: nil,
			},
			args{
				r: audit.Result{
					"tempFolder": "/tmp",
				},
			},
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PhpCompat{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
				resultPath:    tt.fields.resultPath,
				resultFile:    tt.fields.resultFile,
			}
			gotFilename, gotPath, err := p.summaryPath(tt.args.r, tt.args.fileSuffix)
			if (err != nil) != tt.wantErr {
				t.Errorf("PhpCompat.summaryPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFilename != tt.wantFilename {
				t.Errorf("PhpCompat.summaryPath() gotFilename = %v, want %v", gotFilename, tt.wantFilename)
			}
			if gotPath != tt.wantPath {
				t.Errorf("PhpCompat.summaryPath() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestPhpCompat_uploadResults(t *testing.T) {
	type fields struct {
		Report        io.Reader
		ParentProcess audit.Processor
		resultPath    string
		resultFile    string
	}
	type args struct {
		r audit.Result
	}

	// Most coverage is done via Process() tests.

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults *tide.AuditResult
		wantErr     bool
	}{
		{
			"Result Path is empty",
			fields{
				resultPath: "",
				resultFile: "filename.json",
			},
			args{},
			nil,
			true,
		},
		{
			"Result File is empty",
			fields{
				resultPath: "/tmp",
				resultFile: "",
			},
			args{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PhpCompat{
				Report:        tt.fields.Report,
				ParentProcess: tt.fields.ParentProcess,
				resultPath:    tt.fields.resultPath,
				resultFile:    tt.fields.resultFile,
			}
			gotResults, err := p.uploadResults(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PhpCompat.uploadResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("PhpCompat.uploadResults() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
