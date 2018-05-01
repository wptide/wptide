package process

import (
	"github.com/wptide/pkg/process"
	"encoding/json"
	"log"
	"bytes"
	"os"
)

type Debug struct {
	process.Process              // Assume all functions of Process.
	In  <-chan process.Processor // The Processor to intercept.
	Out chan process.Processor   // The outbound Processor.
	N   string                   // A nicename for the process.
}

func (bug *Debug) Run(errc *chan error) error {

	b := bytes.Buffer{}
	log.SetOutput(&b)

	go func() {
		for {
			select {
			case in := <-bug.In:
				bug.CopyFields(in)

				log.SetOutput(os.Stdout)
				log.Printf("DEBUG: %s (%s)\n", bug.N, bug.Message.Title)

				msgJson, _ := json.Marshal(bug.GetResult())

				log.Println(string(msgJson))
				log.SetOutput(&b)

				if bug.Out != nil {
					bug.Out <- bug
				}
			}
		}
	}()

	return nil
}
