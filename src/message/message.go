package message

type Message struct {
	ResponseAPIEndpoint string  `json:"response_api_endpoint"`
	Title               string  `json:"title"`
	Content             string  `json:"content"`
	Slug                string  `json:"slug"`
	ProjectType         string  `json:"project_type,omitempty"`
	SourceURL           string  `json:"source_url"`
	SourceType          string  `json:"source_type"`
	RequestClient       string  `json:"request_client"`
	Force               bool    `json:"force"`
	Visibility          string  `json:"visibility"`
	ExternalRef         *string `json:"external_ref,omitempty"`
}

type MessageProvider interface {
	SendMessage(msg *Message) error
	GetNextMessage() (*Message, error)
	DeleteMessage(ref *string) error
}
