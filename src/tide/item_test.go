package tide

import (
	"reflect"
	"testing"
)

var (
	simpleDummyTheme = &InfoDetailsSimple{
		Name:        "Dummy Theme",
		Version:     "0.0.1",
		Description: "This is a dummy theme.",
	}

	complexDummyTheme = []InfoDetails{
		{
			"Name",
			"Dummy Theme",
		},
		{
			"Version",
			"0.0.1",
		},
		{
			"Description",
			"This is a dummy theme.",
		},
	}
)

func TestSimplifyCodeDetails(t *testing.T) {
	type args struct {
		details []InfoDetails
	}
	tests := []struct {
		name string
		args args
		want *InfoDetailsSimple
	}{
		{
			"Detailed to Simple",
			args{
				complexDummyTheme,
			},
			simpleDummyTheme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyCodeDetails(tt.args.details); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimplifyCodeDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplexifyCodeDetails(t *testing.T) {
	type args struct {
		simple *InfoDetailsSimple
	}
	tests := []struct {
		name string
		args args
		want []InfoDetails
	}{
		{
			"Simple to Complex",
			args{
				simpleDummyTheme,
			},
			complexDummyTheme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComplexifyCodeDetails(tt.args.simple); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComplexifyCodeDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
