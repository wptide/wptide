package phpcompat

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/blang/semver"
	"sort"
)

var (
	phpVersions = map[string]map[string]string{
		"5.2": {
			"max": "5.2.17",
		},
		"5.3": {
			"max": "5.3.29",
		},
		"5.4": {
			"max": "5.4.45",
		},
		"5.5": {
			"max": "5.5.37",
		},
		"5.6": {
			"max": "5.6.32",
		},
		"7.0": {
			"max": "7.0.26",
		},
		"7.1": {
			"max": "7.1.13",
		},
		"7.2": {
			"max": "7.2.1",
		},
	}
	PhpLatest = "7.2.1"
)

type Compatibility struct {
	Source string              `json:"source"`
	Breaks *CompatibilityRange `json:"breaks,omitempty"`
	Warns  *CompatibilityRange `json:"warns,omitempty"`
}

type CompatibilityRange struct {
	Low        string `json:"low"`
	High       string `json:"high"`
	Reported   string `json:"reported"`
	MajorMinor string `json:"major_minor"`
}

func PreviousVersion(version string) string {

	if version == "all" {
		return version
	}

	// Only supporting down to 5.2.0.
	if version == "5.2" || version == "5.2.0" {
		return "5.2.0"
	}

	parts := strings.Split(version, ".")
	if len(parts) < 3 {
		parts = append(parts, "0")
	}

	var maxPrev []string

	if parts[1] == "0" {
		var mMPre string
		switch parts[0] + "." + parts[1] {
		case "7.0":
			mMPre = "5.6"
		}
		maxPrev = strings.Split(phpVersions[mMPre]["max"], ".")
	} else {
		pre, _ := strconv.Atoi(parts[1])
		pre -= 1
		maxPrev = strings.Split(phpVersions[parts[0]+"."+strconv.Itoa(pre)]["max"], ".")
	}

	// Convert and subtract parts
	p3, _ := strconv.Atoi(parts[2])
	p3 -= 1
	if p3 < 0 {
		parts[2] = maxPrev[2]

		if parts[1] == "0" {
			parts[1] = maxPrev[1]
			parts[0] = maxPrev[0]
		} else {

			p2, _ := strconv.Atoi(parts[1])
			p2 -= 1
			parts[1] = strconv.Itoa(p2)

		}
	} else {
		parts[2] = strconv.Itoa(p3)
	}

	return strings.Join(parts, ".")
}

func VersionParts(version string) (int, int, int) {
	if version == "all" {
		return 0, 0, 0
	}
	parts := strings.Split(version, ".")

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	return major, minor, patch
}

func GetVersionParts(version, lowIn string) (low, high, majorMinor, reported string) {
	vParts := strings.Split(version, ".")

	majorMinor = "all"
	high = ""

	// Version "all" is not helpful.
	if len(vParts) > 1 {
		majorMinor = fmt.Sprintf("%s.%s", vParts[0], vParts[1])
	}

	// Is it a Major.Minor? Then get the max.
	if len(vParts) != 1 && len(vParts) != 3 {
		high = phpVersions[majorMinor]["max"]
	} else {
		high = version
	}

	if lowIn == "" {
		if majorMinor != "all" {
			low = fmt.Sprintf("%s.%s", majorMinor, "0")
		} else {
			low = "all"
		}

	} else {
		low = lowIn
	}

	reported = version

	return
}

func BreaksVersions(code string) []string {

	compat := PhpCompatibilityMap[code]

	broken := []string{}

	if compat.Breaks == nil {
		return nil
	}

	var rangeString string
	if compat.Breaks.Reported == "all" {
		rangeString = ">=5.2.0" + " <=" + PhpLatest
	} else {
		rangeString = ">=" + compat.Breaks.Low + " <=" + compat.Breaks.High
	}

	failRange, _ := semver.ParseRange(rangeString)

	for majorMinor, item := range phpVersions {
		if failRange(semver.MustParse(item["max"])) {
			broken = append(broken, majorMinor)
		}
	}

	sort.Strings(broken)

	return broken
}