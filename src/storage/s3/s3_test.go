package s3

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws"
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"os"
)

type mockS3 struct {
	s3manageriface.UploaderAPI
	s3manageriface.DownloaderAPI
}

func (m mockS3) Upload(input *s3manager.UploadInput, _ ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	switch *input.Bucket {
	case "error_bucket":
		return nil, errors.New("something went wrong")
	default:
		return nil, nil
	}
}

func (m mockS3) UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return nil, nil
}

func (m mockS3) Download(_ io.WriterAt, input *s3.GetObjectInput, _ ...func(*s3manager.Downloader)) (int64, error) {
	switch *input.Key {
	case "bucket_error.txt":
		return 0, errors.New("something went wrong")
	default:
		return 0, nil
	}
}

func (m mockS3) DownloadWithContext(aws.Context, io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return 0, nil
}

func mockFileOpen(name string) (*os.File, error) {
	switch name {
	case "error.txt":
		return nil, errors.New("something went wrong")
	default:
		return nil, nil
	}
}

func mockFileCreate(name string) (*os.File, error) {
	switch name {
	case "error.txt":
		return nil, errors.New("something went wrong")
	default:
		return nil, nil
	}
}

func TestS3Provider_UploadFile(t *testing.T) {

	// Set out fileOpen variable to the mock function.
	fileOpen = mockFileOpen
	// Remember to set it back after the test.
	defer func() { fileOpen = os.Open }()

	type args struct {
		filename  string
		reference string
	}
	tests := []struct {
		name    string
		s3p     S3Provider
		args    args
		wantErr bool
	}{
		{
			"Test Upload - upload.txt",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "test_bucket",
			},
			args{
				"upload.txt",
				"upload.txt",
			},
			false,
		},
		{
			"Test Upload Bucket Error",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "error_bucket",
			},
			args{
				"upload.txt",
				"upload.txt",
			},
			true,
		},
		{
			"Test File Open Error",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "test_bucket",
			},
			args{
				"error.txt",
				"error.txt",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s3p.UploadFile(tt.args.filename, tt.args.reference); (err != nil) != tt.wantErr {
				t.Errorf("S3Provider.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestS3Provider_DownloadFile(t *testing.T) {

	// Set out fileOpen variable to the mock function.
	fileCreate = mockFileCreate
	// Remember to set it back after the test.
	defer func() { fileCreate = os.Create }()

	type args struct {
		reference string
		filename  string
	}
	tests := []struct {
		name    string
		s3p     S3Provider
		args    args
		wantErr bool
	}{
		{
			"Test Download - download.txt",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "test_bucket",
			},
			args{
				"download.txt",
				"download.txt",
			},
			false,
		},
		{
			"Test File Creare Error",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "test_bucket",
			},
			args{
				"download.txt",
				"error.txt",
			},
			true,
		},
		{
			"Test Bucket File Error",
			S3Provider{
				uploader:   &mockS3{},
				downloader: &mockS3{},
				bucket:     "test_bucket",
			},
			args{
				"bucket_error.txt",
				"bucket_error.txt",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s3p.DownloadFile(tt.args.reference, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("S3Provider.DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewS3Provider(t *testing.T) {
	type args struct {
		region string
		key    string
		secret string
		bucket string
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			"Test S3 Provider",
			args{
				"us-west-2",
				"random-key",
				"so-secret",
				"the-bucket",
			},
			reflect.TypeOf(&S3Provider{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewS3Provider(tt.args.region, tt.args.key, tt.args.secret, tt.args.bucket); !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("NewS3Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}
