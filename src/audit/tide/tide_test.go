package tide

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"os"
)

var (
	TideClient tide.ClientInterface
	FailClient tide.ClientInterface
)

type MockClient struct {
	apiError bool
}

func (m MockClient) Authenticate(clientId, clientSecret, authEndpoint string) error {
	return nil
}

func (m MockClient) SendPayload(method, endpoint, data string) (string, error) {

	if m.apiError {
		return "", errors.New("API error")
	}

	return "", nil
}

func TestProcessor_Process(t *testing.T) {

	// Clean up after.
	defer func() {
		os.Remove("./testdata/test.json")
	}()

	TideClient = &MockClient{}
	FailClient = &MockClient{apiError: true}

	fullCollectionMessage := message.Message{
		ResponseAPIEndpoint: "http://example/api/tide/v1/audit",
		Title:               "Dummy Theme",
		Content:             "Dummy theme for testing",
		Slug:                "dummy-theme",
		ProjectType:         "theme",
		SourceURL:           "http://themerepo/dummy-theme-1.3.zip",
		SourceType:          "zip",
		RequestClient:       "wporg",
		Force:               false,
		Visibility:          "public",
	}

	fullItemMessage := fullCollectionMessage
	fullItemMessage.ResponseAPIEndpoint += "/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e"

	fullResult := *getFullResult()
	errorResult := *getFullResult()
	errorResult["lighthouseError"] = errors.New("Something went wrong.")
	clientError := *getFullResult()
	clientError["client"] = nil
	noAuditsError := *getFullResult()
	noAuditsError["audits"] = nil
	noExistingResults := *getFullResult()
	noExistingResults["tideItem"] = nil
	apiWriteFail := *getFullResult()
	apiWriteFail["client"] = &FailClient
	lighthouseFail := *getFullResult()
	lighthouseFail["lighthouse"] = nil

	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name    string
		p       Processor
		args    args
		wantErr bool
	}{
		{
			"Collection Audit - Full",
			Processor{},
			args{
				fullCollectionMessage,
				&fullResult,
			},
			false,
		},
		{
			"Item Audit - Full",
			Processor{},
			args{
				fullItemMessage,
				&fullResult,
			},
			false,
		},
		{
			"Lighthouse Error",
			Processor{},
			args{
				fullCollectionMessage,
				&errorResult,
			},
			false,
		},
		{
			"Lighthouse Shell Fail",
			Processor{},
			args{
				fullCollectionMessage,
				&lighthouseFail,
			},
			true,
		},
		{
			"Tide Client Error",
			Processor{},
			args{
				fullCollectionMessage,
				&clientError,
			},
			true,
		},
		{
			"No Audits Error",
			Processor{},
			args{
				fullCollectionMessage,
				&noAuditsError,
			},
			true,
		},
		{
			"Collection Audit - No existing results",
			Processor{},
			args{
				fullCollectionMessage,
				&noExistingResults,
			},
			false,
		},
		{
			"API Write Fail",
			Processor{},
			args{
				fullCollectionMessage,
				&apiWriteFail,
			},
			true,
		},
		{
			"Collection Audit - Save to File",
			Processor{ OutputFile: "./testdata/test.json" },
			args{
				fullCollectionMessage,
				&fullResult,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)

			r := *tt.args.result

			if (r["tideError"] != nil && !tt.wantErr) || (r["tideError"] == nil && tt.wantErr) {
				t.Errorf("Process() error = %v, wantErr %v", r["tideError"], tt.wantErr)
				return
			}

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

func getFullResult() *audit.Result {
	return &audit.Result{
		"client": &TideClient,
		"audits": []string{
			"lighthouse",
		},
		"tempFolder": "/tmp",

		// Added by "ingest" process.
		"ingest":      &Processor{},
		"ingestError": nil,
		"tideItem": &tide.Item{
			Title:       "Dummy Plugin",
			Version:     "1.3",
			Checksum:    "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
			Visibility:  "public",
			ProjectType: "theme",
			SourceUrl:   "http://themerepo/dummy-theme-1.3.zip",
			SourceType:  "zip",
		},
		"checksum": "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
		"files":    []string{"file1.txt", "file2.txt", "file3.txt"},

		// Added by "lighthouse" process.
		"lighthouse": &tide.AuditResult{
			Summary: &tide.AuditSummary{
				LighthouseSummary: &tide.LighthouseSummary{
					ReportCategories: []tide.LighthouseCategory{
						{
							Name:        "Performance",
							Description: "These encapsulate your web app's current performance and opportunities to improve it.",
							Id:          "performance",
							Score:       72.17647058823529,
						},
						{
							Name:        "Progressive Web App",
							Description: "These checks validate the aspects of a Progressive Web App, as specified by the baseline [PWA Checklist](https://developers.google.com/web/progressive-web-apps/checklist).",
							Id:          "pwa",
							Score:       72.17647058823529,
						},
						{
							Name:        "Accessibility",
							Description: "These checks highlight opportunities to [improve the accessibility of your web app](https://developers.google.com/web/fundamentals/accessibility). Only a subset of accessibility issues can be automatically detected so manual testing is also encouraged.",
							Id:          "accessibility",
							Score:       100.00,
						},
						{
							Name:        "Best Practices",
							Description: "We've compiled some recommendations for modernizing your web app and avoiding performance pitfalls.",
							Id:          "best-practices",
							Score:       81.25,
						},
						{
							Name:        "SEO",
							Description: "These checks ensure that your page is optimized for search engine results ranking. There are additional factors Lighthouse does not check that may affect your search ranking. [Learn more](https://support.google.com/webmasters/answer/35769).",
							Id:          "seo",
							Score:       90,
						},
					},
				},
			},
		},
		"lighthouseError": nil,
	}
}

func Test_fallbackValue(t *testing.T) {
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"Fallback - Strings",
			args{
				[]interface{}{
					"",
					"",
					"The Value",
				},
			},
			"The Value",
		},
		{
			"Fallback - Int",
			args{
				[]interface{}{
					0,
					2,
					1,
				},
			},
			2,
		},
		{
			"Fallback - CodeInfo",
			args{
				[]interface{}{
					tide.CodeInfo{},
					tide.CodeInfo{},
					tide.CodeInfo{Type: "theme"},
				},
			},
			tide.CodeInfo{
				Type: "theme",
			},
		},
		{
			"Fallback - Missmatch",
			args{
				[]interface{}{
					"1",
					2,
					&tide.CodeInfo{},
				},
			},
			nil,
		},
		{
			// Note: Custom fallback conditions need to be created.
			// Failure to do so will always return the first item.
			"Fallback - Custom",
			args{
				[]interface{}{
					struct{Name string}{"Rheinard"},
					struct{Name string}{"Someone"},
				},
			},
			struct{Name string}{"Rheinard"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fallbackValue(tt.args.value...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fallbackValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValidItem(t *testing.T) {

	type args struct {
		isCollection bool
		item         tide.Item
		msg          message.Message
		result       audit.Result
		codeInfo     tide.CodeInfo
	}
	tests := []struct {
		name    string
		args    args
		want    PayloadMessage
		wantErr bool
	}{
		{
			"Valid Item",
			args{
				true,
				tide.Item{},
				message.Message{
					RequestClient:"audit_server",
				},
				audit.Result{},
				tide.CodeInfo{},
			},
			PayloadMessage{
				Item:          tide.Item{CodeInfo: tide.CodeInfo{}},
				RequestClient: "audit_server",
			},
			false,
		},
		{
			"Result Code Info",
			args{
				true,
				tide.Item{
					CodeInfo: tide.CodeInfo{Type: "theme"},
				},
				message.Message{
					RequestClient:"audit_server",
				},
				audit.Result{
					"tideItem": &tide.Item{
						CodeInfo: tide.CodeInfo{Type: "other"},
					},
				},
				tide.CodeInfo{
					Type: "theme",
				},
			},
			PayloadMessage{
				Item: tide.Item{
					ProjectType: "other",
					CodeInfo:    tide.CodeInfo{Type: "other"},
				},
				RequestClient: "audit_server",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getValidItem(tt.args.isCollection, tt.args.item, tt.args.msg, tt.args.result, tt.args.codeInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("getValidItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("getValidItem() = %v, want %v", *got, tt.want)
			}
		})
	}
}