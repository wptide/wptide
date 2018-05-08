package payload

import (
	"github.com/wptide/pkg/message"
	"io/ioutil"
	"github.com/wptide/pkg/payload"
	"log"
)

type FilePayload struct{
	TerminateChannel chan struct{}
}

// SendPayload sends the results to a file.
func (fp FilePayload) SendPayload(destination string, payload []byte) ([]byte, error) {
	err := ioutil.WriteFile(destination, payload, 0664)
	if err != nil {
		return nil, err
	}
	if fp.TerminateChannel != nil {
		log.Println("Payload sent to: " + destination)
		// Terminate the server.
		fp.TerminateChannel <- struct{}{}
	}
	return []byte("ok"), nil
}

// BuildPayload uses the default TidePayload.
func (fp FilePayload) BuildPayload(msg message.Message, data map[string]interface{}) ([]byte, error) {
	pl := payload.TidePayload{}
	return pl.BuildPayload(msg, data)
}
