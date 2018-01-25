package info

import (
	"reflect"
	"testing"

	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"github.com/xwp/go-tide/src/audit/ingest"
)

func TestProcessor_Kind(t *testing.T) {
	t.Run("Process Kind", func(t *testing.T) {
		p := Processor{}
		if got := p.Kind(); got != "info" {
			t.Errorf("Processor.Kind() = %v, Impossible, this should be info.", got)
		}
	})
}

var(
	themePath = "./testdata/theme"
	themeResult = &audit.Result{
		"ingest": &ingest.Processor{
			Dest: themePath,
		},
	}

	pluginPath = "./testdata/plugin"
	pluginResult = &audit.Result{
		"ingest": &ingest.Processor{
			Dest: pluginPath,
		},
	}

	otherPath = "./testdata/other"
	otherResult = &audit.Result{
		"ingest": &ingest.Processor{
			Dest: otherPath,
		},
	}

	unknownResult = &audit.Result{
		"ingest": &ingest.Processor{
			Dest: "./unknown",
		},
	}
)

func TestProcessor_Process(t *testing.T) {

	type args struct {
		msg    message.Message
		result *audit.Result
	}
	tests := []struct {
		name string
		p    Processor
		args args
		wantType string
		wantErr bool
	}{
		{
			"Theme Test",
			Processor{},
			args{
				message.Message{},
				themeResult,
			},
			"theme",
			false,
		},
		{
			"Plugin Test",
			Processor{},
			args{
				message.Message{},
				pluginResult,
			},
			"plugin",
			false,
		},
		{
			"Other Code Test",
			Processor{},
			args{
				message.Message{},
				otherResult,
			},
			"other",
			false,
		},
		{
			"Expected Theme",
			Processor{},
			args{
				message.Message{},
				otherResult,
			},
			"theme",
			true,
		},
		{
			"Unknown Path",
			Processor{},
			args{
				message.Message{},
				unknownResult,
			},
			"",
			true,
		},
		{
			"No Ingest Object",
			Processor{},
			args{
				message.Message{},
				&audit.Result{},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process(tt.args.msg, tt.args.result)
			r := *tt.args.result

			// Want type.
			if codeInfo, ok := r["info"].(*tide.CodeInfo); ok {
				if codeInfo.Type != tt.wantType && ! tt.wantErr {
					t.Errorf("Code type: got = %v, want %v", codeInfo.Type, tt.wantType)
				}
			}
		})
	}
}

// Note: Indirectly also covers extractHeader().
func Test_getProjectDetails(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantType    string
		wantDetails []tide.InfoDetails
		wantErr bool
	}{
		{
			"Theme Test",
			args{
				themePath + "/unzipped",
			},
			"theme",
			[]tide.InfoDetails{
				{
					"Description",
					"This is a theme for testing purposes only.",
				},
				{
					"Version",
					"1.0",
				},
				{
					"Author",
					"DummyThemes",
				},
				{
					"AuthorURI",
					"http://dummy.local/",
				},
				{
					"TextDomain",
					"dummy-theme",
				},
				{
					"License",
					"GNU General Public License v2 or later",
				},
				{
					"LicenseURI",
					"http://www.gnu.org/licenses/gpl-2.0.html",
				},
				{
					"Name",
					"Dummy Theme",
				},
				{
					"ThemeURI",
					"http://dummy.local/dummy-theme",
				},
				{
					"Tags",
					"black, brown, orange, tan, white, yellow, light, one-column, two-columns, right-sidebar, flexible-width, custom-header, custom-menu, editor-style, featured-images, microformats, post-formats, rtl-language-support, sticky-post, translation-ready",
				},
			},
			false,
		},
		{
			"Plugin Test",
			args{
				pluginPath + "/unzipped",
			},
			"plugin",
			[]tide.InfoDetails{
				{
					"Name",
					"Dummy Plugin",
				},
				{
					"PluginURI",
					"http://dummy.local/plugin/dummy-plugin",
				},
				{
					"Description",
					"This does nothing.",
				},
				{
					"Version",
					"0.1-alpha",
				},
				{
					"Author",
					"DummyPlugins",
				},
				{
					"AuthorURI",
					"http://dummy.local",
				},
				{
					"TextDomain",
					"dummy-plugin",
				},
				{
					"License",
					"GPL2",
				},
				{
					"LicenseURI",
					"http://www.gnu.org/licenses/gpl-2.0.html",
				},
			},
			false,
		},
		{
			"Other Test",
			args{
				otherPath + "/unzipped",
			},
			"other",
			make([]tide.InfoDetails,0),
			true,
		},
		{
			"Unknown Path",
			args{
				"./unknown",
			},
			"",
			make([]tide.InfoDetails,0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotDetails, err := getProjectDetails(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantType {
				t.Errorf("getProjectDetails() got = %v, want %v", got, tt.wantType)
			}
			if !reflect.DeepEqual(gotDetails, tt.wantDetails) && len(gotDetails) != len(tt.wantDetails) {
				t.Errorf("getProjectDetails() gotDetails = %v, want %v", gotDetails, tt.wantDetails)
			}
		})
	}
}

// Covered by other tests, but need to cover unknown path.
func Test_getCloc(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]tide.ClocResult
		wantErr bool
	}{
		{
			"Unknown Path",
			args{
				"./unknown",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCloc(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCloc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCloc() = %v, want %v", got, tt.want)
			}
		})
	}
}