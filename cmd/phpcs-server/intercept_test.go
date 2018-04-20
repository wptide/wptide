package main

import (
	"testing"

	"github.com/wptide/pkg/process"
	"context"
	"time"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
)

func TestIntercept_Run(t *testing.T) {

	//oldStandard := PHPCompatibilityWPPath
	//PHPCompatibilityWPPath = "/mock/path/to/standard"
	//defer func() {
	//	PHPCompatibilityWPPath = oldStandard
	//}()

	// Need to test with a context.
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	tests := []struct {
		name      string
		procs     []process.Processor
		index int
		wantErrc  bool
		wantErr   bool
		wantValue string
	}{
		{
			name: "Invalid process",
		},
		{
			"PHPCompatibility Theme",
			[]process.Processor{
				&process.Response{
					Process: process.Process{
						Message: message.Message{
							Title:       "Test Theme",
							PayloadType: "mock",
							Audits: &[]message.Audit{
								{"phpcs",
									&message.AuditOption{
										Standard: "phpcompatibility",
									},
								},
							},
						},
						Result: map[string]interface{}{
							"info": &tide.CodeInfo{
								Type: "theme",
							},
						},
					},
				},
			},
			0,
			false,
			false,
			PHPCompatibilityWPPath,
		},
		{
			"PHPCompatibility Plugin 2 Audits",
			[]process.Processor{
				&process.Response{
					Process: process.Process{
						Message: message.Message{
							Title:       "Test Plugin",
							PayloadType: "mock",
							Audits: &[]message.Audit{
								{"phpcs",
									&message.AuditOption{
										Standard: "wordpress",
									},
								},
								{"phpcs",
									&message.AuditOption{
										Standard: "phpcompatibility",
									},
								},
							},
						},
						Result: map[string]interface{}{
							"info": &tide.CodeInfo{
								Type: "plugin",
							},
						},
					},
				},
			},
			1,
			false,
			false,
			PHPCompatibilityWPPath,
		},
		{
			"PHPCompatibility Other",
			[]process.Processor{
				&process.Response{
					Process: process.Process{
						Message: message.Message{
							Title:       "Test Other",
							PayloadType: "mock",
							Audits: &[]message.Audit{
								{"phpcs",
									&message.AuditOption{
										Standard: "phpcompatibility",
									},
								},
							},
						},
						Result: map[string]interface{}{
							"info": &tide.CodeInfo{
								Type: "other",
							},
						},
					},
				},
			},
			0,
			false,
			false,
			"phpcompatibility",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ix := Intercept{}

			ix.SetContext(ctx)
			if tt.procs != nil {
				ix.In = generateProcs(ctx, tt.procs)
			}

			var errc <-chan error
			var err error

			go func() {
				errc, err = ix.Run()
			}()

			// Sleep a short time delay to give process time to start.
			time.Sleep(time.Millisecond * 100)

			if ix.Message.Audits != nil {
				audits := *ix.Message.Audits
				if audits != nil && len(audits) > 0 {
					gotValue := audits[tt.index].Options.Standard
					if gotValue != tt.wantValue {
						t.Errorf("Intercept.Run() gotValue = %v, wantValue %v", gotValue, tt.wantValue)
						return
					}
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Intercept.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (len(errc) != 0) != tt.wantErrc {
				e := <-errc
				t.Errorf("Intercept.Run() error = %v, wantErrc %v", e, tt.wantErrc)
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