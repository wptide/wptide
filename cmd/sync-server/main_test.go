package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
	"errors"
	"github.com/wptide/pkg/message"
	"github.com/xwp/go-tide/src/sync"
	"github.com/xwp/go-tide/src/wporg"
	"os"
)

var mockThemesApi = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := ioutil.ReadAll(r.Body)
	parts := strings.Split(string(bodyByte), "&")

	for _, part := range parts {
		switch part {
		case "request[page]=1":
			file, _ := os.Open("./testdata/themes_request1.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=2":
			file, _ := os.Open("./testdata/themes_request2.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=3":
			file, _ := os.Open("./testdata/themes_request3.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=4":
			fmt.Fprintln(w, `{}`)
			return
		case "request[browse]=fail":
			panic("failed!")
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
		case "request[page]=1":
			file, _ := os.Open("./testdata/plugins_request1.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=2":
			file, _ := os.Open("./testdata/plugins_request2.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=3":
			file, _ := os.Open("./testdata/plugins_request3.json")
			defer file.Close()
			io.Copy(w, file)
			return
		case "request[page]=4":
			fmt.Fprintln(w, `{}`)
			return
		}
	}
	fmt.Fprintln(w, `invalid`)
}))

type mockDispatcher struct{}

func (m mockDispatcher) Dispatch(project wporg.RepoProject) error {
	if project.Slug == "error" {
		return errors.New("something went wrong")
	}
	return nil
}

func (m mockDispatcher) Init() error { return nil }

type mockChecker struct{}

func (m mockChecker) UpdateCheck(project wporg.RepoProject) bool   { return true }
func (m mockChecker) RecordUpdate(project wporg.RepoProject) error { return nil }

func Test_fetcher(t *testing.T) {

	b := bytes.Buffer{}
	log.SetOutput(&b)

	token := make(chan struct{})

	testBufferSize := 5

	type args struct {
		projectType string
		category    string
		bufferSize  int
		token       chan struct{}
		maxPages    int
	}
	type sources struct {
		plugins string
		themes  string
	}
	tests := []struct {
		name    string
		args    args
		sources sources
		want    reflect.Type
	}{
		{
			"Themes",
			args{
				"themes",
				"updated",
				testBufferSize,
				token,
				-1,
			},
			sources{
				mockPluginsApi.URL,
				mockThemesApi.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
		},
		{
			"Plugins - All",
			args{
				"plugins",
				"updated",
				testBufferSize,
				token,
				-1,
			},
			sources{
				mockPluginsApi.URL,
				mockThemesApi.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
		},
		{
			"Plugins - Max Pages 2",
			args{
				"plugins",
				"updated",
				testBufferSize,
				token,
				2,
			},
			sources{
				mockPluginsApi.URL,
				mockThemesApi.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
		},
		{
			"Themes - Fail",
			args{
				"themes",
				"fail",
				testBufferSize,
				token,
				-1,
			},
			sources{
				mockPluginsApi.URL,
				mockThemesApi.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set mock sources.
			client.SetPluginApiSource(tt.sources.plugins)
			client.SetThemeApiSource(tt.sources.themes)

			// Initialise a token.
			go func() { token <- struct{}{} }()

			projects := fetcher(tt.args.projectType, tt.args.category, tt.args.bufferSize, tt.args.token, tt.args.maxPages)

			if reflect.TypeOf(projects) != tt.want {
				t.Errorf("fetcher() = %v, want %v", reflect.TypeOf(projects), tt.want)
			}

			// To avoid deadlock, don't read from the channel when testing fail condition.
			if tt.args.category == "fail" {
				return
			}

			// In the tests we have 3 pages with 5 results. We try to stick to this as it gives enough range
			// for the tests. With maxPages set we'll alter our loop a bit.
			expected := 15
			if tt.args.maxPages > -1 && tt.args.maxPages < 4 {
				expected = tt.args.maxPages * testBufferSize
			}
			for i := 0; i < expected; i++ {
				<-projects
			}
		})
	}
}

func Test_infoWorker(t *testing.T) {

	// Don't send log to os.Stdout.
	b := bytes.Buffer{}
	log.SetOutput(&b)

	type args struct {
		project    wporg.RepoProject
		checker    sync.UpdateChecker
		dispatcher sync.Dispatcher
		messages   chan *message.Message
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Worker - Success",
			args{
				wporg.RepoProject{},
				mockChecker{},
				mockDispatcher{},
				make(chan *message.Message),
			},
		},
		{
			"Worker - Dispatch Error",
			args{
				wporg.RepoProject{
					Slug: "error",
				},
				mockChecker{},
				mockDispatcher{},
				make(chan *message.Message),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projects := make(chan wporg.RepoProject, 1)
			projects <- tt.args.project
			go infoWorker(projects, tt.args.checker, tt.args.dispatcher, tt.args.messages)
			time.Sleep(time.Millisecond * 300)
		})
	}
}

func Test_pool(t *testing.T) {

	// Basic test here, the tests for infoWorker will have
	// the bulk of the testing.

	type args struct {
		workers    int
		project    wporg.RepoProject
		checker    sync.UpdateChecker
		dispatcher sync.Dispatcher
		messages   chan *message.Message
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Pool - Normal",
			args{
				1,
				wporg.RepoProject{},
				mockChecker{},
				mockDispatcher{},
				make(chan *message.Message),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projects := make(chan wporg.RepoProject, 1)
			projects <- tt.args.project
			pool(tt.args.workers, projects, tt.args.checker, tt.args.dispatcher, tt.args.messages)
		})
	}
}

func Test_main(t *testing.T) {

	// Don't log to os.Stdout.
	b := bytes.Buffer{}
	log.SetOutput(&b)

	type serverConfig struct {
		serverActive   bool
		perPage        int
		browseCategory string
		poolSize       int
		quit           chan struct{}
	}

	tests := []struct {
		name         string
		serverConfig serverConfig
	}{
		{
			"Run Main - Active",
			serverConfig{
				true,
				5,
				"updated",
				5,
				make(chan struct{}, 1),
			},
		},
		{
			"Run Main - Not Active",
			serverConfig{
				false,
				5,
				"updated",
				5,
				make(chan struct{}, 1),
			},
		},
	}
	for _, tt := range tests {

		serverActiveOld := serverActive
		serverActive = tt.serverConfig.serverActive
		perPageOld := perPage
		perPage = tt.serverConfig.perPage
		browseCategoryOld := browseCategory
		browseCategory = tt.serverConfig.browseCategory
		poolSizeOld := poolSize
		poolSize = tt.serverConfig.poolSize
		quitOld := quit
		quit = tt.serverConfig.quit

		go func() {
			time.Sleep(time.Second)
			quit <- struct{}{}
		}()

		t.Run(tt.name, func(t *testing.T) {
			main()
		})

		serverActive = serverActiveOld
		perPage = perPageOld
		browseCategory = browseCategoryOld
		poolSize = poolSizeOld
		quit = quitOld
	}
}
