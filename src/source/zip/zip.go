package zip

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"path/filepath"
	"strings"
	"crypto/sha256"
	"archive/zip"
	"sort"
	"encoding/json"
)

type Zip struct {
	url   string
	dest  string
	files []string
	checksum string
}

func (m *Zip) PrepareFiles(dest string) error {

	// Prepare destination.
	m.dest = dest
	if _, err := os.Stat(m.dest); os.IsNotExist(err) {
		os.Mkdir(m.dest, os.ModePerm)
	}

	err := downloadFile(m.url, m.dest + "/source.zip")
	if err !=nil {
		fmt.Println(err)
		return err
	}

	var checksums []string
	m.files, checksums, err = unzip(m.dest+"/source.zip", m.dest+"/unzipped")
	if err !=nil {
		fmt.Println(err)
		return err
	}

	// Calculate checksum - uses same technique as Tide Audit Server.
	sort.Strings(checksums)
	jsonChecksums, _ := json.Marshal(checksums)
	m.checksum = fmt.Sprintf("%x", sha256.Sum256(jsonChecksums))

	return nil
}

func NewZip(url string) *Zip {
	return &Zip{
		url: url,
	}
}

// downloadFile uses an HTTP request to get a file and save it to a given destination folder.
func downloadFile(source string, destination string) error {

	// Create destination
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get file
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

// unzip will un-compress a zip archive,
// moving all files and folders to a destination directory
//
// Props to https://golangcode.com/unzip-files-in-go/ and
// http://blog.ralch.com/tutorial/golang-working-with-zip/
func unzip(source, destination string) (filenames, checksums []string, err error) {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return filenames, checksums, err
	}

	if err := os.MkdirAll(destination, 0755); err != nil {
		return filenames, checksums, err
	}

	rootPath := ""
	for _, file := range reader.File {
		path := file.Name
		if ! file.FileInfo().IsDir() {
			continue
		}
		if len(path) < len(rootPath) || rootPath == "" {
			rootPath = path
		}
	}

	for _, file := range reader.File {
		path := filepath.Join(destination, strings.TrimPrefix(file.Name, rootPath))
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		filenames = append(filenames, path)

		fileReader, err := file.Open()
		if err != nil {
			return filenames, checksums, err
		}
		defer fileReader.Close()
		h := sha256.New()
		if _, err := io.Copy(h, fileReader); err != nil {
			return filenames, checksums, err
		}
		checksums = append(checksums, fmt.Sprintf("%x", h.Sum(nil)))

		fileReader, err = file.Open()
		if err != nil {
			return filenames, checksums, err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return filenames, checksums, err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return filenames, checksums, err
		}
	}

	return filenames, checksums, err
}