package process

import (
	"testing"

	"github.com/wptide/pkg/process"
	"time"
	"context"
	"github.com/wptide/pkg/message"
	"log"
	"os"
	"bytes"
)

func TestDebug_Run(t *testing.T) {

	// Need to test with a context.
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	b := bytes.Buffer{}
	log.SetOutput(&b)

	defer func() {
		log.SetOutput(os.Stdout)
	}()

	tests := []struct {
		name     string
		procs    []process.Processor
		wantErrc bool
		wantErr  bool
	}{
		{
			name: "No input",
		},
		{
			"Simple intercept.",
			[]process.Processor{
				&process.Response{
					Process: process.Process{
						Message: message.Message{
							Title:       "Test Plugin",
							PayloadType: "mock",
							ExternalRef: &[]string{"success"}[0],
						},
						Result: map[string]interface{}{
							"responseSuccess": true,
							"responseMessage": "It worked.",
						},
					},
				},
			},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bug := &Debug{
				Out:     make(chan process.Processor),
			}

			bug.SetContext(ctx)
			if tt.procs != nil {
				bug.In = generateProcs(ctx, tt.procs)
			}

			var err error
			var chanError error
			errc := make(chan error)

			go func() {
				for {
					select {
					case e := <-errc:
						chanError = e
					}
				}
			}()

			go func() {
				err = bug.Run(&errc)
			}()

			// Sleep a short time delay to give process time to start.
			time.Sleep(time.Millisecond * 100)

			if (err != nil) != tt.wantErr {
				t.Errorf("Sink.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (chanError != nil) != tt.wantErrc {
				t.Errorf("Debug.Run() errorChan = %v, wantErrc %v", chanError, tt.wantErrc)
			}
		})
	}
}
