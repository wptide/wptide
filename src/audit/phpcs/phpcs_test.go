package phpcs

import (
	"reflect"
	"testing"

	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
	"github.com/pkg/errors"
	"bytes"
	"github.com/wptide/pkg/log"
	"github.com/wptide/pkg/storage"
	"os"
	"github.com/wptide/pkg/audit/ingest"
	"io"
	"io/ioutil"
)

var (
	cases = map[string]*Processor{
		"undefined": &Processor{},
		"wordpress-core": &Processor{
			Standard: "WordPress-Core",
		},
		"phpcompatibility": &Processor{
			Standard: "PHPCompatibility",
		},
	}

	tempWriter = bytes.Buffer{}

	storageProvider storage.StorageProvider

	// Used by PostProcessor tests
	testFiles = map[string]string{
		"cf7": "./testdata/samples/contactform7-75-phpcs_phpcompatibility.json",
		"t17": "./testdata/samples/twentyseventeen-14-phpcs_wordpress.json",
		"fail": "./testdata/nothing/here",
		"invalid": "./testdata/samples/invalid-json.json",
		"phpcompat": "./testdata/samples/all-phpcompat-sniffs.json",
	}
)

type mockProcessor struct {
	Report        io.Reader
	ParentProcess audit.Processor
}

func (m mockProcessor) Kind() string {
	return "mock"
}

// Mocks a post processor.
// Also serves as an example for building a post-processor.
func (m *mockProcessor) Process(msg message.Message, result *audit.Result) {
	r := *result

	b, err := ioutil.ReadAll(m.Report)
	if err != nil {
		return
	}

	emptyResult := tide.AuditResult{}
	auditResults := r[m.ParentProcess.Kind()].(*tide.AuditResult)

	if auditResults.Summary == emptyResult.Summary {
		auditResults.Summary = &tide.AuditSummary{}
	} else {
		auditResults.Error = string(b)
	}
}

func (m *mockProcessor) SetReport(report io.Reader) {
	m.Report = report
}

func (m *mockProcessor) Parent(parent audit.Processor) {
	m.ParentProcess = parent
}

type mockStorage struct{}

func (m mockStorage) Kind() string {
	return "mock"
}

func (m mockStorage) CollectionRef() string {
	return "mock-collection"
}

func (m mockStorage) UploadFile(filename, reference string) error {
	return nil
}

func (m mockStorage) DownloadFile(reference, filename string) error {
	return nil
}

func TestProcessor_Kind(t *testing.T) {
	tests := []struct {
		name string
		p    *Processor
		want string
	}{
		{
			"Undefined",
			cases["undefined"],
			"phpcs_undefined",
		},
		{
			"WordPress-Core",
			cases["wordpress-core"],
			"phpcs_wordpress-core",
		},
		{
			"PHPCompatibility",
			cases["phpcompatibility"],
			"phpcs_phpcompatibility",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Kind(); got != tt.want {
				t.Errorf("Processor.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_Process(t *testing.T) {

	currentDir, _ := os.Getwd()
	storageProvider = &mockStorage{}

	// Clean up after.
	os.Mkdir("./testdata/tmp/", os.ModePerm)
	defer func() {
		os.RemoveAll("./testdata/tmp/")
	}()

	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name    string
		p       *Processor
		args    args
		hideLog bool
	}{
		{
			"PHPCS Audit",
			&Processor{
				Standard: "WordPress-Core",
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_wordpress-core",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/plugin",
					},
					"checksum":   "27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"tempFolder": currentDir + "/testdata/tmp",
					"fileStore":  &storageProvider,
				},
			},
			true,
		},
		{
			"PHPCS Audit With RuntimeSet",
			&Processor{
				Standard:   "PHPCompatibility",
				RuntimeSet: []string{"testVersion 5.2-"},
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_phpcompatibility",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/plugin",
					},
					"checksum":   "27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"tempFolder": currentDir + "/testdata/tmp",
					"fileStore":  &storageProvider,
				},
			},
			true,
		},
		{
			"PHPCS Audit With PostProcess",
			&Processor{
				Standard:   "PHPCompatibility",
				RuntimeSet: []string{"testVersion 5.2-"},
				PostProcessors: []audit.Processor{
					&mockProcessor{},
				},
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_phpcompatibility",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/plugin",
					},
					"checksum":   "27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"tempFolder": currentDir + "/testdata/tmp",
					"fileStore":  &storageProvider,
				},
			},
			true,
		},
		{
			"PHPCS Audit With Real PostProcess",
			&Processor{
				Standard: "WordPress",
				PostProcessors: []audit.Processor{
					&PhpcsSummary{},
				},
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_wordpress",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/plugin",
					},
					"checksum":   "27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"tempFolder": currentDir + "/testdata/tmp",
					"fileStore":  &storageProvider,
				},
			},
			true,
		},
		{
			"Error: Invalid standard.",
			&Processor{
				Standard: "Invalid",
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_invalid",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/does_not_exist",
					},
					"checksum":   "27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"tempFolder": currentDir + "/testdata/tmp",
					"fileStore":  &storageProvider,
				},
			},
			true,
		},
		{
			"Error: No checksum",
			&Processor{
				Standard: "WordPress-Core",
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_wordpress-core",
					},
					"ingest": &ingest.Processor{
						Dest: "./testdata/unzipped",
					},
				},
			},
			true,
		},
		{
			"Error: No Ingest",
			&Processor{
				Standard: "WordPress-Core",
			},
			args{
				message.Message{},
				&audit.Result{
					"audits": []string{
						"phpcs_wordpress-core",
					},
				},
			},
			true,
		},
		{
			"Error: Can't run this type of audit",
			&Processor{},
			args{
				message.Message{},
				&audit.Result{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hideLog {
				log.SetOutput(&tempWriter)
			} else {
				log.SetOutput(os.Stdout)
			}
			tt.p.Process(tt.args.msg, tt.args.result)
		})
	}
}

func TestProcessor_errCheck(t *testing.T) {

	log.SetOutput(&tempWriter)

	type args struct {
		err    error
		result *audit.Result
	}
	tests := []struct {
		name string
		p    *Processor
		args args
	}{
		{
			"Has Error",
			cases["undefined"],
			args{
				errors.New("this has an error"),
				&audit.Result{},
			},
		},
		{
			"No Error",
			cases["undefined"],
			args{
				nil,
				&audit.Result{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.errCheck(tt.args.err, tt.args.result)
		})
	}
}

func TestProcessor_uploadToStorage(t *testing.T) {

	storageProvider = &mockStorage{}
	log.SetOutput(&tempWriter)

	type args struct {
		r *audit.Result
	}
	tests := []struct {
		name    string
		p       *Processor
		args    args
		want    *tide.AuditResult
		wantErr bool
	}{
		{
			"Upload - Success",
			cases["wordpress-core"],
			args{
				&audit.Result{
					"fileStore":  &storageProvider,
					"tempFolder": "/tmp",
					"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
			&tide.AuditResult{
				Full: struct {
					Type       string `json:"type,omitempty"`
					Key        string `json:"key,omitempty"`
					BucketName string `json:"bucket_name,omitempty"`
				}{
					Type:       "mock",
					Key:        "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_wordpress-core-full.json",
					BucketName: "mock-collection",
				},
			},
			false,
		},
		{
			"Upload - Error storage provider",
			cases["wordpress-core"],
			args{
				&audit.Result{
					"tempFolder": "/tmp",
					"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
				},
			},
			nil,
			true,
		},
		{
			"Upload - Error file path",
			&Processor{
				Standard: "WordPress-Core",
			},
			args{
				&audit.Result{
					"fileStore": &storageProvider,
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.uploadToStorage(*tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Processor.uploadToStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processor.uploadToStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_reportPath(t *testing.T) {

	result := &audit.Result{
		"tempFolder": "/tmp",
		"checksum":   "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
	}

	resultNoTemp := &audit.Result{
		"checksum": "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
	}

	resultNoChecksum := &audit.Result{
		"tempFolder": "/tmp",
	}

	type args struct {
		r *audit.Result
	}
	tests := []struct {
		name         string
		p            *Processor
		args         args
		wantPath     string
		wantFilename string
		wantErr      bool
	}{
		{
			"WordPress-Core",
			cases["wordpress-core"],
			args{
				result,
			},
			"/tmp/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_wordpress-core-full.json",
			"39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_wordpress-core-full.json",
			false,
		},
		{
			"PHPCompatibility",
			cases["phpcompatibility"],
			args{
				result,
			},
			"/tmp/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_phpcompatibility-full.json",
			"39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e-phpcs_phpcompatibility-full.json",
			false,
		},
		{
			"No Temp Folder",
			cases["wordpress-core"],
			args{
				resultNoTemp,
			},
			"",
			"",
			true,
		},
		{
			"No Checksum",
			cases["wordpress-core"],
			args{
				resultNoChecksum,
			},
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.p.reportPath(*tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Processor.reportPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantPath {
				t.Errorf("Processor.reportPath() got = %v, want %v", got, tt.wantPath)
			}
			if got1 != tt.wantFilename {
				t.Errorf("Processor.reportPath() got1 = %v, want %v", got1, tt.wantFilename)
			}
		})
	}
}
