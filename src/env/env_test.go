package env_test

import (
	"testing"
	"os"
	"github.com/xwp/go-tide/src/env"
)

func TestGetEnv(t *testing.T) {

	// Environment variables to set.
	keys := []struct {
		key   string
		value string
	}{
		{
			key:   "ENV_TEST_KEY",
			value: "Provided Value",
		},
	}

	// Range over the keys and remember the original value if it already exists.
	var value string
	for _, ks := range keys {

		// Key is set, so retain the value when the test is finished.
		if value = os.Getenv(ks.key); value != "" {
			os.Unsetenv(ks.key)
			defer func() { os.Setenv(ks.key, value) }()
		}

		// Set the test value.
		os.Setenv(ks.key, ks.value)
	}

	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Variable Provided",
			args: args{
				key:      "ENV_TEST_KEY",
				fallback: "",
			},
			want: "Provided Value",
		},
		{
			name: "Variable Not Provided",
			args: args{
				key:      "ENV_TEST_KEY_EMPTY",
				fallback: "Fallback Value",
			},
			want: "Fallback Value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := env.GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
