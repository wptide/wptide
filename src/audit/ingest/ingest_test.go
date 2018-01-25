package ingest

import (
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"strings"
	"encoding/json"
	"errors"
)

type MockClient struct{}

var (
	TideClient tide.ClientInterface

	mockItems = map[string]*tide.Item{
		"27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a": &tide.Item{},
		"39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e": &tide.Item{
			Results: map[string]tide.AuditResult{
				"lighthouse": tide.AuditResult{
					Summary: &tide.AuditSummary{
						LighthouseSummary: &tide.LighthouseSummary{
							[]tide.LighthouseCategory{
								tide.LighthouseCategory{
									Name:        "Performance",
									Description: "These encapsulate your web app's current performance and opportunities to improve it.",
									Id:          "performance",
									Score:       72.17647058823529,
								},
								tide.LighthouseCategory{
									Name:        "Progressive Web App",
									Description: "These checks validate the aspects of a Progressive Web App, as specified by the baseline [PWA Checklist](https://developers.google.com/web/progressive-web-apps/checklist).",
									Id:          "pwa",
									Score:       72.17647058823529,
								},
								tide.LighthouseCategory{
									Name:        "Accessibility",
									Description: "These checks highlight opportunities to [improve the accessibility of your web app](https://developers.google.com/web/fundamentals/accessibility). Only a subset of accessibility issues can be automatically detected so manual testing is also encouraged.",
									Id:          "accessibility",
									Score:       100.00,
								},
								tide.LighthouseCategory{
									Name:        "Best Practices",
									Description: "We've compiled some recommendations for modernizing your web app and avoiding performance pitfalls.",
									Id:          "best-practices",
									Score:       81.25,
								},
								tide.LighthouseCategory{
									Name:        "SEO",
									Description: "These checks ensure that your page is optimized for search engine results ranking. There are additional factors Lighthouse does not check that may affect your search ranking. [Learn more](https://support.google.com/webmasters/answer/35769).",
									Id:          "seo",
									Score:       90,
								},
							},
						},
					},
				},
			},
		},
	}
)

func (m MockClient) Authenticate(clientId, clientSecret, authEndpoint string) error {
	return nil
}

func (m MockClient) SendPayload(method, endpoint, data string) (string, error) {
	parts := strings.Split(strings.Replace(endpoint, "http://example/api/tide/v1/", "", -1), "/")
	route := parts[0]
	parts = parts[1:]

	switch route {
	case "audit":
		if len(parts) > 0 {
			// Item endpoint.
			if parts[0] == "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" {
				return "bad results", nil
			}

			if i, ok := mockItems[parts[0]]; ok {
				itemBytes, _ := json.Marshal(i)
				return string(itemBytes), nil
			} else {
				return "", errors.New("item not found")
			}
		}
	}

	return "", nil
}

type mockSourceManager struct {
	checksum string
	files    []string
}

func (m *mockSourceManager) PrepareFiles(dest string) error {
	if strings.Split(dest, "/")[1] == "collection" {
		m.checksum = "39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e"
	}
	return nil
}

func (m mockSourceManager) GetChecksum() string {
	return m.checksum
}

func (m mockSourceManager) GetFiles() []string {
	return m.files
}

func TestProcessor_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := Processor{}
		if got := p.Kind(); got != "ingest" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be ingest.", got)
		}
	})
}

func TestProcessor_Process(t *testing.T) {

	TideClient = &MockClient{}

	initialResult := &audit.Result{
		"client": &TideClient,
		"audits": []string{
			"lighthouse",
		},
		"tempFolder": "/tmp",
	}

	initialResultNoTemp := &audit.Result{
		"client": &TideClient,
		"audits": []string{
			"lighthouse",
		},
	}

	initialResultNoAudits := &audit.Result{
		"client":     &TideClient,
		"tempFolder": "/tmp",
	}

	initialResultCollectionResults := &audit.Result{
		"client":     &TideClient,
		"tempFolder": "/collection",
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
		mockSrc bool
	}{
		{
			"Collection Endpoint",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit",
					SourceURL:           "http://example/collection.zip",
				},
				initialResult,
			},
			false,
			true,
		},
		{
			"Collection Endpoint - With Results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit",
					SourceURL:           "http://example/collection.zip",
				},
				initialResultCollectionResults,
			},
			false,
			true,
		},
		{
			"Checksum Endpoint - No results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					SourceURL:           "http://example/noresult.zip",
				},
				initialResult,
			},
			false,
			true,
		},
		{
			"Checksum Endpoint - Lighthouse results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
					SourceURL:           "http://example/result.zip",
				},
				initialResult,
			},
			true,
			true,
		},
		{
			"No Temp Folder",
			&Processor{},
			args{
				message.Message{},
				initialResultNoTemp,
			},
			true,
			true,
		},
		{
			"No Defined Audits",
			&Processor{},
			args{
				message.Message{},
				initialResultNoAudits,
			},
			true,
			true,
		},
		{
			"Bad source URL",
			&Processor{},
			args{
				message.Message{
					SourceURL: "http://example/result.rar",
				},
				initialResult,
			},
			true,
			false,
		},
		{
			"Bad results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				},
				initialResult,
			},
			true,
			true,
		},
		{
			"PrepareFile Error",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					SourceURL:           "http://example/result.zip",
				},
				initialResult,
			},
			true, // A fake url will throw an error so this is expected.
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock the source manager.
			if tt.mockSrc {
				srcMgr = &mockSourceManager{}
			}

			tt.p.Process(tt.args.msg, tt.args.result)

			r := *tt.args.result
			if r["ingestError"] != nil && ! tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", r["ingestError"], tt.wantErr)
				return
			}

			// Reset source manager.
			srcMgr = nil
		})
	}
}

func TestProcessor_getResults(t *testing.T) {

	TideClient = &MockClient{}

	expected := map[string]string{}

	for key, item := range mockItems {
		itemBytes, _ := json.Marshal(item)
		expected[key] = string(itemBytes)
	}

	type args struct {
		client   *tide.ClientInterface
		endpoint string
	}

	tests := []struct {
		name    string
		p       Processor
		args    args
		want    string
		wantErr bool
	}{
		{
			"Checksum Endpoint - Lighthouse results",
			Processor{},
			args{
				&TideClient,
				"http://example/api/tide/v1/audit/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
			},
			expected["39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e"],
			false,
		},
		{
			"Checksum Endpoint - No results",
			Processor{},
			args{
				&TideClient,
				"http://example/api/tide/v1/audit/27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
			},
			expected["27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a"],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getResults(tt.args.client, tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Processor.getResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Processor.getResults() = %v, want %v", got, tt.want)
			}
		})
	}
}
