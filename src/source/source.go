package source

import (
	"strings"
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
