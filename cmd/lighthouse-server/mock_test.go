package main

import (
	"os"
	"io"
	"github.com/wptide/pkg/message"
	"net/http/httptest"
	"net/http"
	"errors"
	"context"
)

/** ----------------------------------------------
	Mock Storage Provider
 */

type mockStorage struct{}

func (m mockStorage) Kind() string {
	return "mock"
}

func (m mockStorage) CollectionRef() string {
	return "mock-collection"
}

// UploadFile simulates an upload and saves the file to ./testdata/upload/{reference}.
func (m mockStorage) UploadFile(filename, reference string) error {

	//switch reference {
	//case "phpcompatuploaderror-phpcs_phpcompatibility-details.json":
	//	fallthrough
	//case "uploaderrorchecksum-phpcs_wordpress-full.json":
	//	return errors.NewPipe("Upload error!")
	//}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	out, _ := os.Create("./testdata/upload/" + reference)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

func (m mockStorage) DownloadFile(reference, filename string) error {
	return nil
}

type mockPayloader struct{}

func (m mockPayloader) BuildPayload(msg message.Message, data map[string]interface{}) ([]byte, error) {

	//if msg.Slug == "buildFail" {
	//	return nil, errors.NewPipe("something went wrong")
	//}

	payload := `{ "valid":"payload" }`
	return []byte(payload), nil
}

func (m mockPayloader) SendPayload(destination string, payload []byte) ([]byte, error) {

	//if destination == "http://test.local/sendfail" {
	//	return nil, errors.NewPipe("something went wrong")
	//}

	reply := `{ "status": "ok" }`
	return []byte(reply), nil
}

/** ----------------------------------------------
	Mock Message Provider
 */

type mockMessageProvider struct {
	Type string
}

func (m mockMessageProvider) SendMessage(msg *message.Message) error {
	return nil
}
func (m mockMessageProvider) GetNextMessage() (*message.Message, error) {
	switch m.Type {
	case "critical":
		pErr := message.NewProviderError("Critical issue.")
		pErr.Type = message.ErrCritcal
		return nil, pErr
	case "quota":
		pErr := message.NewProviderError("Over quota.")
		pErr.Type = message.ErrOverQuota
		return nil, pErr
	case "message":
		return &message.Message{
			Title: "Mock Message",
			ResponseAPIEndpoint: "localhost",

		}, nil
	case "lenError":
		return &message.Message{Title: "lenError"}, nil
	default:
		return nil, nil
	}
}
func (m mockMessageProvider) DeleteMessage(ref *string) error {
	return nil
}

/** ----------------------------------------------
	Mock Tide Client
 */

type mockTide struct {
	apiError bool
}

func (m mockTide) Authenticate(clientId, clientSecret, authEndpoint string) error {
	if clientId == "error" {
		return errors.New("something went wrong")
	}

	return nil
}

func (m mockTide) SendPayload(method, endpoint, data string) (string, error) {

	if m.apiError {
		return "", errors.New("API error")
	}

	return "", nil
}

/** ----------------------------------------------
	Mock file server
 */

var testFileServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
}))

/** ----------------------------------------------
	Mock failed process
 */
type mockFailedProcess struct {
	option    string
	shouldErr bool
}

func (m mockFailedProcess) Run(errc *chan error) error {

	switch m.option {
	case "error":
		*errc <- errors.New("something went wrong")
	}

	if m.shouldErr {
		return errors.New("something went wrong with setup")
	}

	return nil
}

func (m mockFailedProcess) SetContext(ctx context.Context)        {}
func (m mockFailedProcess) SetMessage(msg message.Message)        {}
func (m mockFailedProcess) GetMessage() message.Message           { return message.Message{} }
func (m mockFailedProcess) SetResults(res map[string]interface{}) {}
func (m mockFailedProcess) GetResult() map[string]interface{}     { return nil }
func (m mockFailedProcess) SetFilesPath(path string)              {}
func (m mockFailedProcess) GetFilesPath() string                  { return "" }

