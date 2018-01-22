package util

import (
	"regexp"
)

func IsCollectionEndpoint(url string) bool {
	r, _ := regexp.Compile("([a-fA-F0-9]{64}$)|([0-9]+$)")
	results := r.FindStringSubmatch(url)
	return len(results) == 0
}
