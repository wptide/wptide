package storage

type StorageProvider interface{
	Kind() string
	CollectionRef() string
	UploadFile(filename, reference string) error
	DownloadFile(reference, filename string) error
}