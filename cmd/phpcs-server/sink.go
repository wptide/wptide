package main

import (
	"github.com/wptide/pkg/process"
	"log"
	"github.com/wptide/pkg/message"
)

// Sink is the last process in our Pipeline and responsible for cleanup.
type Sink struct{
	process.Process
	In         <-chan process.Processor
	MessageProvider message.MessageProvider
}

func (s Sink) Run() (<-chan error, error) {

	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		for {
			select {
			case in := <-s.In:
				s.CopyFields(in)
				if s.Result["responseSuccess"] != nil && s.Result["responseSuccess"].(bool) {
					// Output message.
					log.Println(s.Result["responseMessage"])

					// Delete message from provider.
					err := s.MessageProvider.DeleteMessage(s.GetMessage().ExternalRef)
					if err != nil {
						errc <- err
					} else {
						log.Println("'" + s.GetMessage().Title + "' removed from message queue.")
					}
				}
			}
		}

	}()

	return errc, nil
}