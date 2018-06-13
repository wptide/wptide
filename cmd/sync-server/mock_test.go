package main

import (
	"errors"

	"github.com/wptide/pkg/message"
)

// Setup a mockProvider.
type mockProvider struct{}

func (m mockProvider) SendMessage(msg *message.Message) error {
	if msg.Slug == "error" {
		return errors.New("something went wrong")
	}
	return nil
}
func (m mockProvider) GetNextMessage() (*message.Message, error) {
	return nil, nil
}
func (m mockProvider) DeleteMessage(ref *string) error {
	return nil
}
func (m mockProvider) Close() error {
	return nil
}
