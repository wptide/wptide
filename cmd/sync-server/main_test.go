package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/sync"
	"github.com/wptide/pkg/sync/firestore"
	"github.com/wptide/pkg/wporg"
)

var mockThemesAPI = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		}
	}
	fmt.Fprintln(w, `invalid`)
}))

var mockPluginsAPI = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func (m mockDispatcher) Close() error { return nil }

type mockChecker struct {
	LastSync *time.Time
}

func (m mockChecker) UpdateCheck(project wporg.RepoProject) bool         { return true }
func (m mockChecker) RecordUpdate(project wporg.RepoProject) error       { return nil }
func (m mockChecker) SetSyncTime(event, projectType string, t time.Time) {}
func (m mockChecker) GetSyncTime(event, projectType string) time.Time {
	if m.LastSync != nil {
		return *m.LastSync
	}
	return time.Now().AddDate(-10, 0, 0)
}

func Test_fetcher(t *testing.T) {

	oldChecker := checker
	checker = mockChecker{}
	defer func() {
		checker = oldChecker
		defer os.RemoveAll("./db")
	}()

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
		checker sync.UpdateChecker
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
				mockPluginsAPI.URL,
				mockThemesAPI.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
			mockChecker{},
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
				mockPluginsAPI.URL,
				mockThemesAPI.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
			mockChecker{},
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
				mockPluginsAPI.URL,
				mockThemesAPI.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
			mockChecker{},
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
				mockPluginsAPI.URL,
				mockThemesAPI.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
			mockChecker{},
		},
		{
			"Themes - Old",
			args{
				"themes",
				"updated",
				testBufferSize,
				token,
				-1,
			},
			sources{
				mockPluginsAPI.URL,
				mockThemesAPI.URL,
			},
			reflect.TypeOf(make(<-chan wporg.RepoProject)),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set mock sources.
			client.SetPluginAPISource(tt.sources.plugins)
			client.SetThemeAPISource(tt.sources.themes)

			if tt.checker != nil {
				oldChecker := checker
				checker = tt.checker
				defer func() {
					checker = oldChecker
				}()
			} else {
				oldChecker := checker
				checker = &mockChecker{
					LastSync: &[]time.Time{time.Now()}[0],
				}
				defer func() {
					checker = oldChecker
				}()
			}

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

			if tt.checker == nil {
				expected = 0
			}

			for i := 0; i < expected; i++ {
				<-projects
			}

			time.Sleep(time.Millisecond * 10)
		})
	}
}

func Test_infoWorker(t *testing.T) {

	oldChecker := checker
	checker = mockChecker{}
	defer func() {
		checker = oldChecker
	}()

	// Don't send log to os.Stdout.
	b := bytes.Buffer{}
	log.SetOutput(&b)

	type args struct {
		project    wporg.RepoProject
		checker    sync.Updater
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

	oldChecker := checker
	checker = mockChecker{}
	defer func() {
		checker = oldChecker
	}()

	// Basic test here, the tests for infoWorker will have
	// the bulk of the testing.

	type args struct {
		workers    int
		project    wporg.RepoProject
		checker    sync.Updater
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

	oldChecker := checker
	checker = mockChecker{}
	defer func() {
		checker = oldChecker
		defer os.RemoveAll("./db")
	}()

	// Don't log to os.Stdout.
	b := bytes.Buffer{}
	log.SetOutput(&b)

	type serverConfig struct {
		serverActive   bool
		perPage        int
		browseCategory string
		poolSize       int
		quit           chan struct{}
		checkerError   error
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
				nil,
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
				nil,
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
		checkerOld := checker
		checker = &mockChecker{}

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
		checker = checkerOld
	}
}

func Test_mainCheckerError(t *testing.T) {
	oldServerActive := serverActive
	serverActive = true
	oldCheckerError := checkerError
	checkerError = errors.New("error")
	defer func() {
		serverActive = oldServerActive
		checkerError = oldCheckerError
	}()

	go func() {
		time.Sleep(time.Second)
		quit <- struct{}{}
	}()

	t.Run("Checker Fail Test", func(t *testing.T) {
		main()
	})
}

func Test_getSyncProvider(t *testing.T) {
	type args struct {
		c map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			"No Provider",
			args{},
			nil,
		},
		{
			"Local Provider",
			args{
				map[string]map[string]string{
					"app": {"syncDBProvider": "local"},
					"local": {
						"dbPath": "./testdata/testdb",
					},
				},
			},
			reflect.TypeOf(&scribbleChecker{}),
		},
		{
			"Firestore Provider",
			args{
				map[string]map[string]string{
					"app": {
						"syncDBProvider": "firestore",
					},
					"gcp": {
						"projectID": "fake-project-id-12345",
						"docPath":   "sync-server/wporg",
					},
				},
			},
			reflect.TypeOf(&firestore.Sync{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getSyncProvider(tt.args.c); reflect.TypeOf(got) != tt.want {
				t.Errorf("getSyncProvider() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}

func Test_getDispatcher(t *testing.T) {
	type args struct {
		c map[string]map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Type
		wantErr bool
	}{
		{
			"Firestore Dispatcher",
			args{
				map[string]map[string]string{
					"app": {
						"syncPHPCSActive":      "on",
						"syncLighthouseActive": "on",
						"messageProvider":      "firestore",
					},
					"gcp": {
						"projectID":            "fake-id",
						"docPath":              "doc",
						"lighthouseCollection": "fake-lh",
						"phpcsCollection":      "fake-phpcs",
					},
				},
			},
			reflect.TypeOf(&firestoreDispatcher{}),
			false,
		},
		{
			"SQS Dispatcher",
			args{
				map[string]map[string]string{
					"app": {
						"syncPHPCSActive":      "on",
						"syncLighthouseActive": "on",
						"messageProvider":      "sqs",
					},
					"aws": {
						"key":             "abc",
						"secret":          "def",
						"sqs_region":      "us-west-2",
						"sqs_lh_queue":    "fake-lh",
						"sqs_phpcs_queue": "fake-phpcs",
					},
				},
			},
			reflect.TypeOf(&sqsDispatcher{}),
			false,
		},
		{
			"MongoDB Dispatcher",
			args{
				map[string]map[string]string{
					"app": {
						"syncPHPCSActive":      "on",
						"syncLighthouseActive": "on",
						"messageProvider":      "local",
					},
					"mongo": {
						"user":                 "",
						"pass":                 "",
						"host":                 "localhost:12345",
						"database":             "queue",
						"lighthouseCollection": "fake-lh",
						"phpcsCollection":      "fake-phpcs",
					},
				},
			},
			reflect.TypeOf(&mongoDispatcher{}),
			false,
		},
		{
			"Nil Dispatcher",
			args{},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDispatcher(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDispatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got) != tt.want {
				t.Errorf("getDispatcher() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
