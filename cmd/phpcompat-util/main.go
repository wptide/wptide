// NOT FOR DISTRIBUTION.
// This program is purely a utility for creating code. It should not be included in any distribution.

package main

import (
	"io"
	"flag"
	"os"
	"log"
	"bytes"
	"github.com/wptide/pkg/tide"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"github.com/wptide/pkg/phpcompat"
)

var (
	verbs = []string{

		//"available since",

		"not present", // or earlier
		////"as of", // use "soft reserved"
		"soft reserved", // "as of" "reserved keyword as of" , works like "deprecated"
		////"introduced", // use "reserved" instead
		"reserved",   // "since", "introduced", "as of"
		"deprecated", // optional "removed"
		////"removed since", // use "removed"
		"removed",
		"removed in", // added since PHP 7 (see what I did there?)
		////"since", // Probably should not use this one.
		"supported",
		////"Found:",
		"prior to",
		"or earlier",
		"and earlier",
		"and lower",
		"<",
		////"in PHP", // Probably don't need this one

		"magic method",

		"available since",

		"since",
	}
)

func main() {

	input := flag.String("input", "", "Please specify the file to use as the source for the map.")
	output := flag.String("output", "", "Please specify the file to use for the map output.")
	flag.Parse()

	var fileReader io.Reader
	var fileWriter io.WriteCloser
	var err error

	if fileReader, err = os.Open(*input); err != nil {
		log.Fatal("Input file could not be opened.")
	}

	if fileWriter, err = os.Create(*output); err != nil {
		log.Fatal("Output file could not be created.")
	}

	defer fileWriter.Close()

	generateCompatibilityMap(fileReader, fileWriter)
}

func generateCompatibilityMap(input io.Reader, output io.Writer) {

	data := bytes.Buffer{}
	io.Copy(&data, input)

	var fullResults *tide.PhpcsResults
	err := json.Unmarshal(data.Bytes(), &fullResults)
	if err != nil {
		log.Fatal(err)
	}

	sources := make(map[string]map[string]string)

	for _, data := range fullResults.Files {
		for _, msg := range data.Messages {
			sources[msg.Source] = make(map[string]string)
			sources[msg.Source]["message"] = msg.Message
			sources[msg.Source]["type"] = msg.Type
		}
	}

	output.Write([]byte("package phpcompat\n\nvar PhpCompatibilityMap = map[string]*Compatibility{\n"))
	for src, item := range sources {
		sourceOutput := formatAsSourceCode(getSourceOutput(src, item["message"], item["type"]))
		//sourceOutput := formatAsSourceCode(getSourceOutput(src, item["message"]))
		if len(sourceOutput) > 0 {
			output.Write([]byte(sourceOutput))

		}
	}
	output.Write([]byte("\n}\n"))
}

func formatAsSourceCode(source string) string {

	var compat phpcompat.Compatibility
	err := json.Unmarshal([]byte(source), &compat)

	if err != nil {
		return ""
	}

	key := compat.Source

	breaksLow := ""
	breaksHigh := ""
	breaksReported := ""
	breaksMajorMinor := ""
	breaksOutput := ""
	if compat.Breaks != nil {
		breaksLow = compat.Breaks.Low
		breaksHigh = compat.Breaks.High
		breaksReported = compat.Breaks.Reported
		breaksMajorMinor = compat.Breaks.MajorMinor
		breaksOutput =
			`
		Breaks: &CompatibilityRange{
			Low: "` + breaksLow + `",
			High: "` + breaksHigh + `",
			Reported: "` + breaksReported + `",
			MajorMinor: "` + breaksMajorMinor + `",
		},`
	}

	warnsLow := ""
	warnsHigh := ""
	warnsReported := ""
	warnsMajorMinor := ""
	warnsOutput := ""
	if compat.Warns != nil {
		warnsLow = compat.Warns.Low
		warnsHigh = compat.Warns.High
		warnsReported = compat.Warns.Reported
		warnsMajorMinor = compat.Warns.MajorMinor
		warnsOutput =
			`
		Warns: &CompatibilityRange{
			Low:        "` + warnsLow + `",
			High:       "` + warnsHigh + `",
			Reported:   "` + warnsReported + `",
			MajorMinor: "` + warnsMajorMinor + `",
		},`
	}

	output := `
	"` + key + `": &Compatibility{
		Source: "` + key + `",` +
		breaksOutput +
		warnsOutput + `
	},`

	return output
}

func getSourceOutput(src, msg, kind string) string {
	line := fmt.Sprintf("%s\n\t%s\n", src, msg)
	//line := fmt.Sprintf("%s\n", msg)

	versions := getVersions(line)
	var breaks *phpcompat.CompatibilityRange
	var warns *phpcompat.CompatibilityRange

	if strings.ToLower(kind) == "error" {
		low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "")
		breaks = &phpcompat.CompatibilityRange{
			low,
			high,
			reported,
			majorMinor,
		}
	} else {
		low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "")
		warns = &phpcompat.CompatibilityRange{
			low,
			high,
			reported,
			majorMinor,
		}
	}

	matchList := strings.Join(verbs, "|")

	var re = regexp.MustCompile(`(?i)` + matchList + ``)
	matches := re.FindAllString(msg, -1)

	if len(matches) > 0 {

		matches = orderMatches(matches)

		// NOTE: Order is VERY important
		switch matches[0] {
		case "supported":
			fallthrough
		case "not present":
			low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "5.2.0")

			breaks = &phpcompat.CompatibilityRange{
				low,
				high,
				reported,
				majorMinor,
			}
		case "soft reserved":
			fallthrough
		case "deprecated":

			low, _, majorMinor, reported := phpcompat.GetVersionParts(versions[0], versions[0])
			warns = &phpcompat.CompatibilityRange{
				low,
				phpcompat.PhpLatest,
				reported,
				majorMinor,
			}

			if len(versions) > 1 {
				low, _, majorMinor, reported := phpcompat.GetVersionParts(versions[1], "")
				breaks = &phpcompat.CompatibilityRange{
					low,
					phpcompat.PhpLatest,
					reported,
					majorMinor,
				}

				warns.High = phpcompat.PreviousVersion(breaks.Low)
			}
		case "reserved":
			// We don't want to pick up "soft reserved".
			if len(matches) == 1 || (len(matches) > 1 && matches[0] != matches[1]) {

				low, _, majorMinor, reported := phpcompat.GetVersionParts(versions[0], versions[0])

				breaks = &phpcompat.CompatibilityRange{
					low,
					phpcompat.PhpLatest,
					reported,
					majorMinor,
				}
			}
		case "removed in":
			fallthrough
		case "removed":
			// We don't want to pick up "soft reserved" or "deprecated".
			if len(versions) == 1 {

				low, _, majorMinor, reported := phpcompat.GetVersionParts(versions[0], versions[0])

				breaks = &phpcompat.CompatibilityRange{
					low,
					phpcompat.PhpLatest,
					reported,
					majorMinor,
				}
			}
		case "prior to":

			// Annoying bad grammar for this error...
			if src == "PHPCompatibility.PHP.EmptyNonVariable.Found" {
				version := phpcompat.PreviousVersion(versions[0])
				low, high, majorMinor, reported := phpcompat.GetVersionParts(version, "5.2.0")

				breaks = &phpcompat.CompatibilityRange{
					low,
					high,
					reported,
					majorMinor,
				}
			} else {
				low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], versions[0])

				breaks = &phpcompat.CompatibilityRange{
					low,
					high,
					reported,
					majorMinor,
				}
			}
		case "<":
			// Another annoying grammar.
			if src == "PHPCompatibility.PHP.TernaryOperators.MiddleMissing" {
				low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], versions[0])

				breaks = &phpcompat.CompatibilityRange{
					low,
					high,
					reported,
					majorMinor,
				}
			}
		case "or earlier":
			fallthrough
		case "and earlier":
			fallthrough
		case "and lower":
			low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "5.2.0")

			breaks = &phpcompat.CompatibilityRange{
				low,
				high,
				reported,
				majorMinor,
			}

		case "magic method":
			low, high, majorMinor, reported := phpcompat.GetVersionParts("all", "")

			breaks = &phpcompat.CompatibilityRange{
				low,
				high,
				reported,
				majorMinor,
			}

		case "available since":
			if breaks != nil {
				breaks.Low = "5.2.1"
				breaks.High = phpcompat.PreviousVersion(breaks.MajorMinor)
			}
			if warns != nil {
				warns.Low = "5.2.1"
				warns.High = phpcompat.PreviousVersion(warns.MajorMinor)
			}

		case "since":
			if breaks != nil {
				breaks.High = phpcompat.PhpLatest
			}
			if warns != nil {
				warns.High = phpcompat.PhpLatest
			}
		}

		if breaks == nil && warns == nil {
			return ""
		}
	}
	report := phpcompat.Compatibility{
		src,
		breaks,
		warns,
	}

	incompatString, _ := json.Marshal(report)

	return string(incompatString)
}

func getSourceOutputByKind(src, msg, kind string) string {
	line := fmt.Sprintf("%s\n\t%s\n", src, msg)

	versions := getVersions(line)
	var breaks *phpcompat.CompatibilityRange
	var warns *phpcompat.CompatibilityRange

	if strings.ToLower(kind) == "error" {
		low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "")
		breaks = &phpcompat.CompatibilityRange{
			low,
			high,
			reported,
			majorMinor,
		}
	} else {
		low, high, majorMinor, reported := phpcompat.GetVersionParts(versions[0], "")
		warns = &phpcompat.CompatibilityRange{
			low,
			high,
			reported,
			majorMinor,
		}
	}

	matchList := strings.Join(verbs, "|")
	var re = regexp.MustCompile(`(?i)` + matchList + ``)
	matches := re.FindAllString(msg, -1)

	if len(matches) > 0 {

	}

	if breaks == nil && warns == nil {
		return ""
	}

	report := phpcompat.Compatibility{
		src,
		breaks,
		warns,
	}

	incompatString, _ := json.Marshal(report)

	return string(incompatString)
}

func getVersions(line string) []string {
	pattern := `(?i)((\d+\.)+\d+)|(\ball\b)`
	var re = regexp.MustCompile(pattern)
	result := re.FindAllString(line, -1)

	if len(result) == 0 {
		// Becaise they don't like minors?
		pattern := `(?i)PHP 7`
		var re = regexp.MustCompile(pattern)
		result := re.FindAllString(line, -1)

		if len(result) == 0 {
			return []string{"all"}
		}

		return []string{"7.0"}
	}

	return result
}

func orderMatches(m []string) (matches []string) {

	found := func(value string, matches []string) bool {
		for _, val := range matches {
			if val == value {
				return true
			}
		}
		return false
	}

	for _, verb := range verbs {

		for _, match := range m {

			if strings.ToLower(verb) == strings.ToLower(match) && ! found(match, matches) {
				matches = append(matches, strings.ToLower(match))
			}
		}
	}

	return;
}
