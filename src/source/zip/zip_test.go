package zip

import (
	"reflect"
	"testing"
)

func Test_combinedChecksum(t *testing.T) {
	type args struct {
		sums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Successful Sum (Sorted)",
			args{
				[]string{
					"27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"2c8b08da5ce60398e1f19af0e5dccc744df274b826abe585eaba68c525434806",
					"f6936912184481f5edd4c304ce27c5a1a827804fc7f329f43d273b8621870776",
				},
			},
			"5a0c0a95d189c266ca1ed43767dd98f3fb513ce3434e2b08f34828ac11e79a94",
		},
		{
			"Successful Sum (Unsorted)",
			args{
				[]string{
					"f6936912184481f5edd4c304ce27c5a1a827804fc7f329f43d273b8621870776",
					"27dd8ed44a83ff94d557f9fd0412ed5a8cbca69ea04922d88c01184a07300a5a",
					"2c8b08da5ce60398e1f19af0e5dccc744df274b826abe585eaba68c525434806",
				},
			},
			"5a0c0a95d189c266ca1ed43767dd98f3fb513ce3434e2b08f34828ac11e79a94",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinedChecksum(tt.args.sums); got != tt.want {
				t.Errorf("combinedChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_GetChecksum(t *testing.T) {

	checksum := "5a0c0a95d189c266ca1ed43767dd98f3fb513ce3434e2b08f34828ac11e79a94"

	// Should be impossible to fail.
	t.Run("Get Checksum", func(t *testing.T) {
		m := Zip{
			checksum: checksum,
		}
		if got := m.GetChecksum(); got != checksum {
			t.Errorf("Zip.GetChecksum() = %v, want %v", got, checksum)
		}
	})
}

func TestZip_GetFiles(t *testing.T) {

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	// Should be impossible to fail.
	t.Run("Get Files", func(t *testing.T) {
		m := Zip{
			files: files,
		}
		if got := m.GetFiles(); !reflect.DeepEqual(got, files) {
			t.Errorf("Zip.GetFiles() = %v, want %v", got, files)
		}
	})
}
