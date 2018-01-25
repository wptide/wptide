package tide

import (
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"github.com/pkg/errors"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)

			r := *tt.args.result

			if (r["tideError"] != nil && ! tt.wantErr) || (r["tideError"] == nil && tt.wantErr) {
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
