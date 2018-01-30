package lighthouse

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/storage"
	"io/ioutil"
	"github.com/pkg/errors"
)

var(
	storageProvider storage.StorageProvider
)

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
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

func mockWriteFile(filename string, data []byte, perm os.FileMode) error {
	switch filename {
	case "/tmp/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff-lighthouse-full.json":
		return errors.New("something went wrong")
	default:
		return nil
	}
}

func TestProcessor_Process(t *testing.T) {

	// Set out execCommand variable to the mock function.
	execCommand = mockExecCommand
	// Remember to set it back after the test.
	defer func() { execCommand = exec.Command }()

	// Set out execCommand variable to the mock function.
	writeFile = mockWriteFile
	// Remember to set it back after the test.
	defer func() { writeFile = ioutil.WriteFile }()

	storageProvider = mockStorage{}

	result := &audit.Result{
		"audits": []string{"lighthouse"},
		"fileStore": &storageProvider,
		"tempFolder": "/tmp",
		"checksum": "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
	}

	noAuditResult := &audit.Result{}

	noFileStoreResult := &audit.Result{
		"audits": []string{"lighthouse"},
	}

	noTempResult := &audit.Result{
		"audits": []string{"lighthouse"},
		"fileStore": &storageProvider,
	}

	noChecksumResult := &audit.Result{
		"audits": []string{"lighthouse"},
		"fileStore": &storageProvider,
		"tempFolder": "/tmp",
	}

	errorWriteResult := &audit.Result{
		"audits": []string{"lighthouse"},
		"fileStore": &storageProvider,
		"tempFolder": "/tmp",
		"checksum": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name    string
		p       *Processor
		args    args
		wantErr bool
	}{
		{
			name: "Successful Report",
			p:    &Processor{},
			args: args{
				msg:    message.Message{},
				result: result,
			},
			wantErr: false,
		},
		{
			name: "JSON Error",
			p:    &Processor{},
			args: args{
				msg:    message.Message{Slug: "jsonError"},
				result: result,
			},
			wantErr: true,
		},
		{
			name: "Lighthouse Command Error",
			p:    &Processor{},
			args: args{
				msg:    message.Message{Slug: "error"},
				result: result,
			},
			wantErr: true,
		},
		{
			name: "Don't Run Lighthouse",
			p: &Processor{},
			args: args{
				msg:    message.Message{},
				result: noAuditResult,
			},
			wantErr: false,
		},
		{
			name: "Write File Error - Upload",
			p:    &Processor{},
			args: args{
				msg:    message.Message{},
				result: errorWriteResult,
			},
			wantErr: true,
		},
		{
			name: "No File Store - Upload",
			p:    &Processor{},
			args: args{
				msg:    message.Message{},
				result: noFileStoreResult,
			},
			wantErr: true,
		},
		{
			name: "No Temp Folder - Upload",
			p:    &Processor{},
			args: args{
				msg:    message.Message{},
				result: noTempResult,
			},
			wantErr: true,
		},
		{
			name: "No Checksum - Upload",
			p:    &Processor{},
			args: args{
				msg:    message.Message{},
				result: noChecksumResult,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)
			r := *tt.args.result
			if (r["lighthouseError"] != nil && !tt.wantErr) ||
				(r["lighthouseError"] == nil && tt.wantErr) {
				t.Errorf("lighthouse.Process: Error %v, WantErr %v", r["lighthouseError"], tt.wantErr)
			}
		})
	}
}

// TestHelperProcess is the fake command.
func TestHelperProcess(t *testing.T) {

	// If the helper process var is not set this code should not run.
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// Exit helper sub routine if nothing else exits.
	defer os.Exit(0)

	// Get the passed arguments.
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}

	// If no arguments, write to Stderr and exit
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]

	switch cmd {

	case "lh":
		switch args[0] {
		case "https://wp-themes.com/error":
			fmt.Fprintf(os.Stderr, "Error occurred.")
			os.Exit(1)
		case "https://wp-themes.com/jsonError":
			fmt.Fprintf(os.Stdout, "Invalid json.")
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stdout, exampleReport())
			os.Exit(0)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}

func exampleReport() string {
	return `{
  "reportCategories": [
    {
      "name": "Performance",
      "description": "These encapsulate your web app's current performance and opportunities to improve it.",
      "id": "performance",
      "score": 72.17647058823529
    },
    {
      "name": "Progressive Web App",
      "weight": 1,
      "description": "These checks validate the aspects of a Progressive Web App, as specified by the baseline [PWA Checklist](https://developers.google.com/web/progressive-web-apps/checklist).",
      "id": "pwa",
      "score": 54.54545454545455
    },
    {
      "name": "Accessibility",
      "description": "These checks highlight opportunities to [improve the accessibility of your web app](https://developers.google.com/web/fundamentals/accessibility). Only a subset of accessibility issues can be automatically detected so manual testing is also encouraged.",
      "id": "accessibility",
      "score": 100
    },
    {
      "name": "Best Practices",
      "description": "We've compiled some recommendations for modernizing your web app and avoiding performance pitfalls.",
      "id": "best-practices",
      "score": 81.25
    },
    {
      "name": "SEO",
      "description": "These checks ensure that your page is optimized for search engine results ranking. There are additional factors Lighthouse does not check that may affect your search ranking. [Learn more](https://support.google.com/webmasters/answer/35769).",
      "id": "seo",
      "score": 90
    }
  ]
}`
}

func TestProcessor_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := Processor{}
		if got := p.Kind(); got != "lighthouse" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be lighthouse.", got)
		}
	})
}
