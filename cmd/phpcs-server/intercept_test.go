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
			"No Code Info",
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
						Result: map[string]interface{}{},
					},
				},
			},
			0,
			false,
			false,
			"",
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
							"info": tide.CodeInfo{
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
							"info": tide.CodeInfo{
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
							"info": tide.CodeInfo{
								Type: "other",
							},
						},
					},
				},
			},
			0,
			false,
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ix := Intercept{}

			ix.SetContext(ctx)
			if tt.procs != nil {
				ix.In = generateProcs(ctx, tt.procs)
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
				err = ix.Run(&errc)
			}()

			// Sleep a short time delay to give process time to start.
			time.Sleep(time.Millisecond * 100)

			if ix.Message.Audits != nil {
				audits := *ix.Message.Audits
				if audits != nil && len(audits) > 0 {
					gotValue := audits[tt.index].Options.StandardOverride
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

			if (chanError != nil) != tt.wantErrc {
				t.Errorf("Intercept.Run() errorChan = %v, wantErrc %v", chanError, tt.wantErrc)
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