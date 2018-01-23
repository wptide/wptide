package source

import (
	"reflect"
	"testing"
)

func TestGetKind(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple Zip",
			args: args{
				url: "http://example.local/example.zip",
			},
			want: "zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKind(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}