package process

import (
	"testing"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/process"
	"time"
	"context"
	"errors"
)

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
	return nil, nil
}
func (m mockMessageProvider) DeleteMessage(ref *string) error {
	if *ref == "fail" {
		return errors.New("something went wrong.")
	}

	return nil
}

func TestSink_Run(t *testing.T) {

	// Need to test with a context.
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	mockProvider := &mockMessageProvider{}

	type fields struct {
		MessageProvider message.MessageProvider
	}
	tests := []struct {
		name     string
		fields   fields
		procs    []process.Processor
		wantErrc bool
		wantErr  bool
	}{
		{
			name: "No Response",
		},
		{
			"Remove from provider - success",
			fields{
				mockProvider,
			},
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
		{
			"Remove from provider - fail",
			fields{
				mockProvider,
			},
			[]process.Processor{
				&process.Response{
					Process: process.Process{
						Message: message.Message{
							Title:       "Test Plugin Fail",
							PayloadType: "mock",
							ExternalRef: &[]string{"fail"}[0],
						},
						Result: map[string]interface{}{
							"responseSuccess": true,
							"responseMessage": "...but message removal failed.",
						},
					},
				},
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sink{
				MessageProvider: tt.fields.MessageProvider,
			}

			s.SetContext(ctx)
			if tt.procs != nil {
				s.In = generateProcs(ctx, tt.procs)
			}

			var errc <-chan error
			var err error

			go func() {
				errc, err = s.Run()
			}()

			// Sleep a short time delay to give process time to start.
			time.Sleep(time.Millisecond * 100)

			if (err != nil) != tt.wantErr {
				t.Errorf("Response.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (len(errc) != 0) != tt.wantErrc {
				e := <-errc
				t.Errorf("Response.Run() error = %v, wantErrc %v", e, tt.wantErrc)
				return
			}
		})
	}
}

func generateProcs(ctx context.Context, procs []process.Processor) <-chan process.Processor {
	out := make(chan process.Processor, len(procs))
	go func() {
		for _, proc := range procs {
			proc.SetContext(ctx)
			out <- proc
		}
	}()
	return out
}