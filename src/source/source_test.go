package source

import (
	"reflect"
	"testing"
	"github.com/xwp/go-tide/src/source/zip"
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

func TestNewSourceManager(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			name: "Simple Zip",
			args: args{
				url: "http://example.local/example.zip",
			},
			want: reflect.TypeOf(&zip.Zip{}),
		},
		{
			name: "Something Else",
			args: args{
				url: "http://example.local/example.tgz",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSourceManager(tt.args.url); !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("NewSourceManager() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
