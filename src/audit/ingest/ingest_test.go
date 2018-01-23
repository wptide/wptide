package ingest

import (
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"fmt"
	"strings"
	"encoding/json"
	"errors"
)

type MockClient struct{}

var (
	TideClient tide.ClientInterface
	mockItems  = map[string]*tide.Item{

		"27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a": &tide.Item{},
		"39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e": &tide.Item{
			Results: map[string]tide.AuditResult{
				"lighthouse": tide.AuditResult{
					Summary: struct {
						tide.PhpcsSummary
						tide.LighthouseSummary
					}{
						tide.PhpcsSummary{},
						tide.LighthouseSummary{
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
			if i, ok := mockItems[parts[0]]; ok {
				itemBytes, _ := json.Marshal(i)
				return string(itemBytes), nil
			} else {
				return "", errors.New("item not found")
			}
		} else {
			// Collection.
			fmt.Println("collection")
		}
	}

	return "", nil
}

func TestProcessor_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := Processor{}
		if got := p.Kind(); got != "lighthouse" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be lighthouse.", got)
		}
	})
}

func TestProcessor_Process(t *testing.T) {

	TideClient = &MockClient{}

	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name string
		p    *Processor
		args args
	}{
		{
			"Collection Endpoint",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit",
					SourceURL: "http://example/collection.zip",
				},
				&audit.Result{
					"client": &TideClient,
					"audits": []string{
						"lighthouse",
					},
				},
			},
		},
		{
			"Checksum Endpoint - No results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					SourceURL: "http://example/noresult.zip",
				},
				&audit.Result{
					"client": &TideClient,
					"audits": []string{
						"lighthouse",
					},
				},
			},
		},
		{
			"Checksum Endpoint - Lighthouse results",
			&Processor{},
			args{
				message.Message{
					ResponseAPIEndpoint: "http://example/api/tide/v1/audit/39c7d71a68565ddd7b6a0fd68d94924d0db449a99541439b3ab8a477c5f1fc4e",
					SourceURL: "http://example/result.zip",
				},
				&audit.Result{
					"client": &TideClient,
					"audits": []string{
						"lighthouse",
					},
				},
			},
		},
		//{
		//	"Checksum Endpoint - No Results"
		//},
		//{
		//	"Checksum Endpoint - Lighthouse Results"
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)
			fmt.Println(tt.args.result)
		})
	}
}

func TestProcessor_getResults(t *testing.T) {
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
		// TODO: Add test cases.
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

func getClient(client *tide.ClientInterface) *tide.ClientInterface {
	return client
}
