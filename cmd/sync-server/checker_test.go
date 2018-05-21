package main

import (
	"os"
	"reflect"
	"testing"

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
				lastUpdated: &wporg.RepoProject{
					LastUpdated: "1970-01-01 12:00am, MST",
				},
			}
			if err := s.RecordUpdate(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("scribbleChecker.RecordUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scribbleChecker_GetLastUpdated(t *testing.T) {
	testDriver, _ := scribble.New("./testdata/testdb", nil)
	notFoundDriver, _ := scribble.New("./testdata/", nil)

	type fields struct {
		db          *scribble.Driver
		lastUpdated *wporg.RepoProject
	}
	tests := []struct {
		name    string
		fields  fields
		want    *wporg.RepoProject
		wantErr bool
	}{
		{
			"DB Error",
			fields{
				db: nil,
			},
			nil,
			true,
		},
		{
			"Get Last Updated",
			fields{
				db: testDriver,
			},
			&wporg.RepoProject{
				Name:        "Theme One",
				Slug:        "theme-one",
				Version:     "1.0.1",
				LastUpdated: "2018-02-28 12:00am GMT",
			},
			false,
		},
		{
			"Get Last Updated - Not Found",
			fields{
				db: notFoundDriver,
			},
			&wporg.RepoProject{
				LastUpdated: "1970-01-01 12:00am MST",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scribbleChecker{
				db:          tt.fields.db,
				lastUpdated: tt.fields.lastUpdated,
			}
			got, err := s.GetLastUpdated()
			if (err != nil) != tt.wantErr {
				t.Errorf("scribbleChecker.GetLastUpdated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scribbleChecker.GetLastUpdated() = %v, want %v", got, tt.want)
			}
		})
	}
}
