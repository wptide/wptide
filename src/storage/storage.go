package storage

type StorageProvider interface{
	UploadFile(filename, reference string) error
	DownloadFile(reference, filename string) error
}