package log

import (
	"testing"
	"bytes"
	"os"
)

func TestLog(t *testing.T) {

	// Use a buffer while testing to avoid showing output in tests.
	tempWriter := bytes.Buffer{}
	logger.SetOutput(&tempWriter)
	// Then set it back to Stdout.
	defer func() { logger.SetOutput(os.Stdout) }()

	testInterface := make(map[string]interface{})
	testInterface["title"] = "Interface Title"
	testInterface["content"] = "Interface Content"

	type args struct {
		title   interface{}
		content interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Integers",
			args{
				1,
				2,
			},
		},
		{
			"Floats",
			args{
				1.2,
				2.4,
			},
		},
		{
			"Strings",
			args{
				"The Title",
				"The Content",
			},
		},
		{
			"Structs",
			args{
				fromStruct().title,
				fromStruct().content,
			},
		},
		{
			"Interfaces",
			args{
				testInterface["title"],
				testInterface["content"],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Log(tt.args.title, tt.args.content)
		})
	}
}

func fromStruct() struct {
	title   string
	content string
} {
	return struct {
		title   string
		content string
	}{
		"Struct Title",
		"Struct Content",
	}
}
