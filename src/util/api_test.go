package util

import (
	"testing"
)

func TestIsCollectionEndpoint(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Collection Endpoint",
			args{
				url: "http://example/api/tide/v1/audit",
			},
			true,
		},
		{
			"Post Endpoint",
			args{
				url: "http://example/api/tide/v1/audit/123",
			},
			false,
		},
		{
			"Checksum Endpoint",
			args{
				url: "http://example/api/tide/v1/audit/4750ac90919d42c8c0f54da61d47180d93c15edf30099157e20e718f0e0ac649",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCollectionEndpoint(tt.args.url); got != tt.want {
				t.Errorf("IsCollectionEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
