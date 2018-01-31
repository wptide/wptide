package zip

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var fileServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	switch r.URL.String() {
	case "/test.zip":
		w.Header().Set("Content-Type", "applicaiton/zip")
		w.Header().Set("Content-Disposition", "attachment; filename='test.zip'")
		http.ServeFile(w, r, "./testdata/test.zip")
	}
}))

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

func Test_unzip(t *testing.T) {
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name          string
		args          args
		wantFilenames []string
		wantChecksums []string
		wantErr       bool
	}{
		{
			"Unzip File - Success",
			args{
				"./testdata/test.zip",
				"./testdata/unzipped",
			},
			[]string{
				"testdata/unzipped/function.php",
				"testdata/unzipped/script.js",
				"testdata/unzipped/style.css",
			},
			[]string{
				"64a43b6ce686b50bbd7eb91b2b1346ed66e7053d42f7f7b9d5562d55a25d1321",
				"9a8549c5d1f384593788dc25b1c236f8450534e8cb95833003786fef8201b92b",
				"09679b8abb88b21dd1cf166e1d2745df7882a879d2b8672548f6dc0dc9572fe6",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFilenames, gotChecksums, err := unzip(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("unzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFilenames, tt.wantFilenames) {
				t.Errorf("unzip() gotFilenames = %v, want %v", gotFilenames, tt.wantFilenames)
			}
			if !reflect.DeepEqual(gotChecksums, tt.wantChecksums) {
				t.Errorf("unzip() gotChecksums = %v, want %v", gotChecksums, tt.wantChecksums)
			}
		})
	}
}

func TestZip_PrepareFiles(t *testing.T) {
	type fields struct {
		url      string
		dest     string
		files    []string
		checksum string
	}
	type args struct {
		dest string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Simple Zip Fetch",
			fields{
				url:  fileServer.URL + "/test.zip",
				dest: "./testdata/download/",
			},
			args{
				dest: "./testdata/download/",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Zip{
				url:      tt.fields.url,
				dest:     tt.fields.dest,
				files:    tt.fields.files,
				checksum: tt.fields.checksum,
			}
			if err := m.PrepareFiles(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Zip.PrepareFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
