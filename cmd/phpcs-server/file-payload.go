package main

import (
	"github.com/wptide/pkg/message"
	"io/ioutil"
	"github.com/wptide/pkg/payload"
)

type filePayload struct{}

// SendPayload sends the results to a file.
func (fp filePayload) SendPayload(destination string, payload []byte) ([]byte, error) {
	err := ioutil.WriteFile(destination, payload, 0664)
	if err != nil {
		return nil, err
	}
	return []byte("ok"), nil
}

// BuildPayload uses the default TidePayload.
func (fp filePayload) BuildPayload(msg message.Message, data map[string]interface{}) ([]byte, error) {
	pl := payload.TidePayload{}
	return pl.BuildPayload(msg, data)
}
