package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/process"
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
		return &message.Message{}, nil
	case "lenError":
		return &message.Message{Title: "lenError"}, nil
	default:
		return nil, nil
	}
}
func (m mockMessageProvider) DeleteMessage(ref *string) error {
	if ref != nil && *ref == "removeFail" {
		return errors.New("something went wrong")
	}
	return nil
}

func (m mockMessageProvider) Close() error {
	return nil
}

/** ----------------------------------------------
Mock Tide Client
*/

type mockTide struct {
	apiError bool
}

func (m mockTide) Authenticate(clientID, clientSecret, authEndpoint string) error {
	if clientID == "error" {
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
Mock phpcs process
*/
type mockPHPCSProcess struct {
	err error
}

func (m mockPHPCSProcess) Run(errc *chan error) error     { return nil }
func (m mockPHPCSProcess) SetContext(ctx context.Context) {}
func (m mockPHPCSProcess) SetMessage(msg message.Message) {}
func (m mockPHPCSProcess) GetMessage() message.Message    { return message.Message{} }
func (m mockPHPCSProcess) SetResults(res *process.Result) {}
func (m mockPHPCSProcess) GetResult() *process.Result     { return nil }
func (m mockPHPCSProcess) SetFilesPath(path string)       {}
func (m mockPHPCSProcess) GetFilesPath() string           { return "" }

func (m mockPHPCSProcess) Do() error {
	if m.err != nil {
		return m.err
	}
	return nil
}

/** ----------------------------------------------
Mock processes
*/

func mockProcResponse(pre string, msg message.Message) error {
	if msg.Slug == pre+"Fail" {
		return errors.New("something went wrong")
	}
	return nil
}

func MockDoProcess(proc process.Processor, msg message.Message, result *process.Result) error {
	switch proc.(type) {
	case *process.Ingest:
		if msg.Slug == "resultSuccess" {
			res := *result
			res["responseSuccess"] = true
			res["responseMessage"] = "Success!"
			result = &res
		}
		return mockProcResponse("ingest", msg)
	case *process.Info:
		return mockProcResponse("info", msg)
	case *process.Phpcs:
		return mockProcResponse("phpcs", msg)
	case *process.Response:
		return mockProcResponse("res", msg)
	}
	return nil
}
