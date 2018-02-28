package wporg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"sort"
)

var mockThemesApi = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := ioutil.ReadAll(r.Body)
	parts := strings.Split(string(bodyByte), "&")

	for _, part := range parts {
		switch part {
		case "request[per_page]=5":
			fmt.Fprintln(w, wantedThemeResponse5Json)
			return
		}
	}
	fmt.Fprintln(w, `invalid`)
}))

var mockPluginsApi = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := ioutil.ReadAll(r.Body)
	parts := strings.Split(string(bodyByte), "&")

	for _, part := range parts {
		switch part {
		case "request[browse]=alternate":
			fmt.Fprintln(w, wantedPluginResponse5JsonAlternate)
			return
		case "request[per_page]=5":
			fmt.Fprintln(w, wantedPluginResponse5Json)
			return
		case "request[browse]=httpError":
			panic("noo!")
		}
	}
	fmt.Fprintln(w, `invalid`)
}))

var
(
	wantedPluginResponse5Json = `{
				"info": {
					"page": 1,
					"pages": 1024,
 					"results": 5
				},
				"plugins": [
					{"name": "Plugin 1", "slug": "plugin-1", "version": "1.0.0", "last_updated": "2018-02-21", "short_description": "Plugin 1 Short", "description": "", "download_link": "http://example.local/plugin-1-1.0.0.zip" },
					{"name": "Plugin 2", "slug": "plugin-2", "version": "2.0.0", "last_updated": "2018-02-22", "short_description": "Plugin 2 Short", "description": "", "download_link": "http://example.local/plugin-2-2.0.0.zip" },
					{"name": "Plugin 3", "slug": "plugin-3", "version": "3.0.0", "last_updated": "2018-02-23", "short_description": "Plugin 3 Short", "description": "", "download_link": "http://example.local/plugin-3-3.0.0.zip" },
					{"name": "Plugin 4", "slug": "plugin-4", "version": "4.0.0", "last_updated": "2018-02-24", "short_description": "Plugin 4 Short", "description": "", "download_link": "http://example.local/plugin-4-4.0.0.zip" },
					{"name": "Plugin 5", "slug": "plugin-5", "version": "5.0.0", "last_updated": "2018-02-25", "short_description": "Plugin 5 Short", "description": "", "download_link": "http://example.local/plugin-5-5.0.0.zip" }
				]
			}`
	wantedPluginResponse5JsonAlternate = `{
				"info": {
					"page": 1,
					"pages": 1024,
 					"results": 5
				},
				"plugins": {
					"plugin-1": {"name": "Plugin 1", "slug": "plugin-1", "version": "1.0.0", "last_updated": "2018-02-21", "short_description": "Plugin 1 Short", "description": "", "download_link": "http://example.local/plugin-1-1.0.0.zip" },
					"plugin-2": {"name": "Plugin 2", "slug": "plugin-2", "version": "2.0.0", "last_updated": "2018-02-22", "short_description": "Plugin 2 Short", "description": "", "download_link": "http://example.local/plugin-2-2.0.0.zip" },
					"plugin-3": {"name": "Plugin 3", "slug": "plugin-3", "version": "3.0.0", "last_updated": "2018-02-23", "short_description": "Plugin 3 Short", "description": "", "download_link": "http://example.local/plugin-3-3.0.0.zip" },
					"plugin-4": {"name": "Plugin 4", "slug": "plugin-4", "version": "4.0.0", "last_updated": "2018-02-24", "short_description": "Plugin 4 Short", "description": "", "download_link": "http://example.local/plugin-4-4.0.0.zip" },
					"plugin-5": {"name": "Plugin 5", "slug": "plugin-5", "version": "5.0.0", "last_updated": "2018-02-25", "short_description": "Plugin 5 Short", "description": "", "download_link": "http://example.local/plugin-5-5.0.0.zip" }
				}
			}`
	wantedPluginResponse5 = ApiResponse{
		Info: ApiInfo{
			1,
			1024,
			5,
		},
		Type: "plugins",
		Plugins: []RepoProject{
			{
				"Plugin 1",
				"plugin-1",
				"1.0.0",
				"2018-02-21",
				"Plugin 1 Short",
				"",
				"http://example.local/plugin-1-1.0.0.zip",
				"",
			},
			{
				"Plugin 2",
				"plugin-2",
				"2.0.0",
				"2018-02-22",
				"Plugin 2 Short",
				"",
				"http://example.local/plugin-2-2.0.0.zip",
				"",
			},
			{
				"Plugin 3",
				"plugin-3",
				"3.0.0",
				"2018-02-23",
				"Plugin 3 Short",
				"",
				"http://example.local/plugin-3-3.0.0.zip",
				"",
			},
			{
				"Plugin 4",
				"plugin-4",
				"4.0.0",
				"2018-02-24",
				"Plugin 4 Short",
				"",
				"http://example.local/plugin-4-4.0.0.zip",
				"",
			},
			{
				"Plugin 5",
				"plugin-5",
				"5.0.0",
				"2018-02-25",
				"Plugin 5 Short",
				"",
				"http://example.local/plugin-5-5.0.0.zip",
				"",
			},
		},
	}
	wantedThemeResponse5Json = `{
				"info": {
					"page": 1,
					"pages": 1024,
 					"results": 5
				},
				"themes": [
					{"name": "Theme 1", "slug": "theme-1", "version": "1.0.0", "last_updated": "2018-02-21", "description": "Theme 1 Short", "short_description": "", "download_link": "http://example.local/theme-1-1.0.0.zip" },
					{"name": "Theme 2", "slug": "theme-2", "version": "2.0.0", "last_updated": "2018-02-22", "description": "Theme 2 Short", "short_description": "", "download_link": "http://example.local/theme-2-2.0.0.zip" },
					{"name": "Theme 3", "slug": "theme-3", "version": "3.0.0", "last_updated": "2018-02-23", "description": "Theme 3 Short", "short_description": "", "download_link": "http://example.local/theme-3-3.0.0.zip" },
					{"name": "Theme 4", "slug": "theme-4", "version": "4.0.0", "last_updated": "2018-02-24", "description": "Theme 4 Short", "short_description": "", "download_link": "http://example.local/theme-4-4.0.0.zip" },
					{"name": "Theme 5", "slug": "theme-5", "version": "5.0.0", "last_updated": "2018-02-25", "description": "Theme 5 Short", "short_description": "", "download_link": "http://example.local/theme-5-5.0.0.zip" }
				]
			}`
	wantedThemeResponse5 = ApiResponse{
		Info: ApiInfo{
			1,
			1024,
			5,
		},
		Type: "themes",
		Themes: []RepoProject{
			{
				"Theme 1",
				"theme-1",
				"1.0.0",
				"2018-02-21",
				"",
				"Theme 1 Short",
				"http://example.local/theme-1-1.0.0.zip",
				"",
			},
			{
				"Theme 2",
				"theme-2",
				"2.0.0",
				"2018-02-22",
				"",
				"Theme 2 Short",
				"http://example.local/theme-2-2.0.0.zip",
				"",
			},
			{
				"Theme 3",
				"theme-3",
				"3.0.0",
				"2018-02-23",
				"",
				"Theme 3 Short",
				"http://example.local/theme-3-3.0.0.zip",
				"",
			},
			{
				"Theme 4",
				"theme-4",
				"4.0.0",
				"2018-02-24",
				"",
				"Theme 4 Short",
				"http://example.local/theme-4-4.0.0.zip",
				"",
			},
			{
				"Theme 5",
				"theme-5",
				"5.0.0",
				"2018-02-25",
				"",
				"Theme 5 Short",
				"http://example.local/theme-5-5.0.0.zip",
				"",
			},
		},
	}
)

func TestClient_Request(t *testing.T) {
	type args struct {
		source      string
		projectType string
		category    string
		perPage     int
		page        int
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    *ApiResponse
		wantErr bool
	}{
		{
			"Get 5 Themes",
			Client{},
			args{
				mockThemesApi.URL,
				"themes",
				"updated",
				5,
				1,
			},
			&wantedThemeResponse5,
			false,
		},
		{
			"Get 5 Plugins",
			Client{},
			args{
				mockPluginsApi.URL,
				"plugins",
				"updated",
				5,
				1,
			},
			&wantedPluginResponse5,
			false,
		},
		{
			"Get 5 Plugins - Alternate Response",
			Client{},
			args{
				mockPluginsApi.URL,
				"plugins",
				"alternate",
				5,
				1,
			},
			&wantedPluginResponse5,
			false,
		},
		{
			"Themes invalid JSON",
			Client{},
			args{
				mockPluginsApi.URL,
				"themes",
				"invalid",
				-1,
				1,
			},
			nil,
			true,
		},
		{
			"Plugin HTTP Request Error",
			Client{},
			args{
				"https://fakeurl",
				"plugins",
				"httpError",
				-1,
				1,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Request(tt.args.source, tt.args.projectType, tt.args.category, tt.args.perPage, tt.args.page)

			// DeepEqual doesn't work for maps converted to slices, so we have to alter this a bit.
			if err == nil && got != nil && got.Plugins != nil && len(got.Plugins) > 1 {
				sort.SliceStable(got.Plugins, func(i, j int) bool { return got.Plugins[i].Name < got.Plugins[j].Name })
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RequestThemes(t *testing.T) {

	// Mock the API.
	oldUrl := themesApiUrl
	themesApiUrl = mockThemesApi.URL
	defer func() {
		themesApiUrl = oldUrl
	}()

	type args struct {
		category string
		perPage  int
		page     int
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    *ApiResponse
		wantErr bool
	}{
		{
			"Get 5 Themes",
			Client{},
			args{
				"updated",
				5,
				1,
			},
			&wantedThemeResponse5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.RequestThemes(tt.args.category, tt.args.perPage, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RequestThemes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RequestThemes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RequestPlugins(t *testing.T) {

	// Mock the API.
	oldUrl := pluginsApiUrl
	pluginsApiUrl = mockPluginsApi.URL
	defer func() {
		pluginsApiUrl = oldUrl
	}()

	type args struct {
		category string
		perPage  int
		page     int
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    *ApiResponse
		wantErr bool
	}{
		{
			"Get 5 Plugins",
			Client{},
			args{
				"updated",
				5,
				1,
			},
			&wantedPluginResponse5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.RequestPlugins(tt.args.category, tt.args.perPage, tt.args.page)

			// DeepEqual doesn't work for maps converted to slices, so we have to alter this a bit.
			if err == nil && len(got.Plugins) > 1 {
				sort.SliceStable(got.Plugins, func(i, j int) bool { return got.Plugins[i].Name < got.Plugins[j].Name })
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RequestPlugins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RequestPlugins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetPluginApiSource(t *testing.T) {
	type fields struct {
		pluginAPI string
	}
	type args struct {
		source string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"Set plugin API url",
			fields{
				"",
			},
			args{
				"http://plugins.local",
			},
			"http://plugins.local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pluginAPI: tt.fields.pluginAPI,
			}
			c.SetPluginApiSource(tt.args.source)
			got := c.pluginAPI
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SetPluginApiSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetThemeApiSource(t *testing.T) {
	type fields struct {
		themeAPI string
	}
	type args struct {
		source string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want string
	}{
		{
			"Set theme API url",
			fields{
				"",
			},
			args{
				"http://themes.local",
			},
			"http://themes.local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				themeAPI: tt.fields.themeAPI,
			}
			c.SetThemeApiSource(tt.args.source)
			got := c.themeAPI
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SetThemeApiSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
