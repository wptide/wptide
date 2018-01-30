package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	fileCreate = os.Create
	fileOpen   = os.Open
)

type S3Provider struct {
	session    *session.Session
	uploader   s3manageriface.UploaderAPI
	downloader s3manageriface.DownloaderAPI
	bucket     string
}

func (s3p S3Provider) Kind() string {
	return "s3"
}

func (s3p S3Provider) CollectionRef() string {
	return s3p.bucket
}

func (s3p S3Provider) UploadFile(filename, reference string) error {

	// Open file for writing to S3.
	file, err := fileOpen(filename)

	// Error if file cannot be opened.
	if err != nil {
		return err
	}
	defer file.Close()

	// Use the upload manager to write to S3.
	_, err = s3p.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3p.bucket),
		Key:    aws.String(reference),
		Body:   file,
	})

	// Error if file cannot be uploaded.
	if err != nil {
		return err

	}

	return nil
}

func (s3p S3Provider) DownloadFile(reference, filename string) error {

	// Create file for writing.
	file, err := fileCreate(filename)

	// Error if file cannot be created.
	if err != nil {
		return err
	}
	defer file.Close()

	// Attempt to download file from S3.
	_, err = s3p.downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3p.bucket),
			Key:    aws.String(reference),
		})

	// Error on failed download.
	if err != nil {
		return err
	}

	return nil
}

// NewS3Provider is a convenience method to return a new *S3Provider instance.
func NewS3Provider(region, key, secret, bucket string) *S3Provider {

	sess, _ := getSession(region, key, secret)
	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)

	return &S3Provider{
		session:    sess,
		uploader:   uploader,
		downloader: downloader,
		bucket:     bucket,
	}
}

// getSession establishes a new SQS session.
func getSession(region, key, secret string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	return sess, err
}
