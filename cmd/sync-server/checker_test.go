package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/nanobox-io/golang-scribble"
	"github.com/wptide/pkg/wporg"
)

var testDriver *scribble.Driver

func Test_scribbleChecker_UpdateCheck(t *testing.T) {

	testDriver, _ := scribble.New("./testdata/testdb", nil)

	type fields struct {
		db *scribble.Driver
	}
	type args struct {
		project wporg.RepoProject
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Check Theme - Needs Update",
			fields{
				testDriver,
			},
			args{
				wporg.RepoProject{
					Name:        "Project One",
					Slug:        "project-one",
					LastUpdated: "2018-02-28",
					Version:     "1.0.1",
					Type:        "themes",
				},
			},
			true,
		},
		{
			"Check Plugin - Up to date",
			fields{
				testDriver,
			},
			args{
				wporg.RepoProject{
					Name:        "Plugin One",
					Slug:        "plugin-one",
					LastUpdated: "2018-02-28",
					Version:     "1.0.0",
					Type:        "plugins",
				},
			},
			false,
		},
		{
			"Check Theme - No Driver",
			fields{
				nil,
			},
			args{
				wporg.RepoProject{
					Name:        "Project One",
					Slug:        "project-one",
					LastUpdated: "2018-02-28",
					Version:     "1.0.1",
					Type:        "themes",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scribbleChecker{
				db: tt.fields.db,
			}
			if got := s.UpdateCheck(tt.args.project); got != tt.want {
				t.Errorf("scribbleChecker.UpdateCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scribbleChecker_RecordUpdate(t *testing.T) {

	testDriver, _ := scribble.New("./testdata/testdb", nil)

	// Clean up after.
	defer func() {
		os.RemoveAll("./testdata/testdb/themes/theme-one.json")
	}()

	type fields struct {
		db *scribble.Driver
	}
	type args struct {
		project wporg.RepoProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Write Theme - Success",
			fields{
				testDriver,
			},
			args{
				wporg.RepoProject{
					Name:        "Theme One",
					Slug:        "theme-one",
					LastUpdated: "2018-02-28 12:00am GMT",
					Version:     "1.0.1",
					Type:        "themes",
				},
			},
			false,
		},
		{
			"Write Theme - Fail",
			fields{
				nil,
			},
			args{
				wporg.RepoProject{
					Name:        "Theme One",
					Slug:        "theme-one",
					LastUpdated: "2018-02-28 12:00am GMT",
					Version:     "1.0.1",
					Type:        "themes",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scribbleChecker{
				db: tt.fields.db,
			}
			if err := s.RecordUpdate(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("scribbleChecker.RecordUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scribbleChecker_GetSyncTime(t *testing.T) {

	testDriver, _ := scribble.New("./testdata/testdb", nil)
	epoc, _ := time.Parse(wporg.TimeFormat, "1970-01-01 12:00am MST")

	type fields struct {
		db *scribble.Driver
	}
	type args struct {
		event       string
		projectType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			"Test Sync Time Error",
			fields{
				testDriver,
			},
			args{
				"start",
				"theme",
			},
			epoc,
		},
		{
			"Get Sync Time",
			fields{
				testDriver,
			},
			args{
				"start",
				"mock",
			},
			time.Unix(0, 1526956796534182580),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scribbleChecker{
				db: tt.fields.db,
			}
			if got := s.GetSyncTime(tt.args.event, tt.args.projectType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scribbleChecker.GetSyncTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scribbleChecker_SetSyncTime(t *testing.T) {

	testDriver, _ := scribble.New("./testdata/testdb", nil)
	epoc, _ := time.Parse(wporg.TimeFormat, "1970-01-01 12:00am MST")

	defer os.Remove("./testdata/testdb/info/mock-sync-stop.json")

	type fields struct {
		db *scribble.Driver
	}
	type args struct {
		event       string
		projectType string
		t           time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"Set Sync Time",
			fields{
				testDriver,
			},
			args{
				"stop",
				"mock",
				epoc,
			},
			"0", // Epoc is 0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scribbleChecker{
				db: tt.fields.db,
			}
			s.SetSyncTime(tt.args.event, tt.args.projectType, tt.args.t)

			b, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/testdb/info/%s-sync-%s.json", tt.args.projectType, tt.args.event))
			if string(b) != tt.want {
				t.Errorf("scribbleChecker.GetSyncTime() = %v, expected %v", string(b), tt.want)
			}

		})
	}
}
