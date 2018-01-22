package source

import (
	"strings"
	"github.com/xwp/go-tide/src/source/zip"
)

type Source interface {
	PrepareFiles(dest string) error
}

func GetKind(url string) string {
	var kind string
	ts := strings.Split(url, ".")
	if len(ts) > 1 {
		kind = ts[len(ts)-1]
	}
	return kind
}

func NewSourceManager(url string) Source {
	kind := GetKind(url)
	switch kind {
	case "zip":
		return zip.NewZip(url)
	default:
		return nil
	}
}
