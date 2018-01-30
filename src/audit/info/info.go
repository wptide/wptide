package info

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/audit/ingest"
	"fmt"
	"github.com/hhatto/gocloc"
	"github.com/pkg/errors"
	"github.com/xwp/go-tide/src/tide"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
	"github.com/xwp/go-tide/src/log"
)

type Processor struct{}

func (p Processor) Kind() string {
	return "info"
}

// Process takes *result and sends relevant results to the Tide API.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	log.Log(msg.Title, "Processing CodeInfo")

	ingestObj, ok := r["ingest"].(*ingest.Processor);
	if ! ok || ingestObj.Dest == "" {
		r["infoError"] = errors.New("could not find path to code")
		r["info"] = nil

		log.Log(msg.Title, r["infoError"])
		return
	}

	path := ingestObj.Dest + "/unzipped"

	cloc, err := getCloc(path)
	if err != nil {
		r["infoError"] = err
		r["info"] = nil

		log.Log(msg.Title, r["infoError"])
		return
	}

	projectType, details, _ := getProjectDetails(path)

	r["info"] = &tide.CodeInfo{
		Type:    projectType,
		Details: details,
		Cloc:    cloc,
	}

	log.Log(msg.Title, "Project is `" + projectType + "`")
}

func getProjectDetails(path string) (string, []tide.InfoDetails, error) {

	projectType := "other"
	details := []tide.InfoDetails{}

	var found bool

	// Get files in root path.
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", nil, err
	}

	// Traverse files and scan for headers.
	for _, f := range files {
		projectType, details, err = extractHeader(path + "/" + f.Name())
		if err == nil {
			found = true
			break
		}
	}

	if ! found {
		err = errors.New("not a theme or plugin")
	}

	return projectType, details, err
}

func getCloc(path string) (map[string]tide.ClocResult, error) {

	clocMap := make(map[string]tide.ClocResult)

	languages := gocloc.NewDefinedLanguages()
	options := gocloc.NewClocOptions()
	paths := []string{path}

	processor := gocloc.NewProcessor(languages, options)
	cloc, err := processor.Analyze(paths)

	if err != nil {
		return nil, err
	}

	clocTotals := tide.ClocResult{0, 0, 0, 0}

	for _, cLang := range cloc.Languages {
		// Add Totals
		clocTotals.Code += int(cLang.Code)
		clocTotals.Blank += int(cLang.Blanks)
		clocTotals.Comment += int(cLang.Comments)
		clocTotals.NFiles += len(cLang.Files)

		clocMap[strings.ToLower(cLang.Name)] = tide.ClocResult{
			int(cLang.Blanks),
			int(cLang.Comments),
			int(cLang.Code),
			len(cLang.Files),
		}
	}

	clocMap["sum"] = clocTotals

	return clocMap, nil
}

// extractHeader scans every .php file in the path to retrieve a possible plugin header, or
// looks for style.css to extract the theme header.
//
// The information is return as an interface map.
func extractHeader(filename string) (projectType string, details []tide.InfoDetails, err error) {

	projectType = "other"

	headerFields := []string{
		"Plugin Name",
		"Plugin URI",
		"Description",
		"Version",
		"Author",
		"Author URI",
		"Text Domain",
		"License",
		"License URI",
		"Theme Name",
		"Theme URI",
		"Tags",
	}

	f, _ := os.Open(filename)
	defer f.Close()
	b1 := make([]byte, 8192)
	n1, _ := f.Read(b1)

	isStyleCSS, _ := regexp.Match(`(style.css)$`, []byte(filename))

	if n1 > 0 {

		validHeader := false
		for _, field := range headerFields {
			pattern := fmt.Sprintf("%s:.*", field)
			re := regexp.MustCompile(pattern)
			value := strings.Replace(re.FindString(string(b1)), field+":", "", -1)
			if len(value) > 0 {

				fieldname := field

				if field == "Plugin Name" {
					projectType = "plugin"
					validHeader = true
					fieldname = "Name"
				}
				if field == "Theme Name" && isStyleCSS {
					projectType = "theme"
					validHeader = true
					fieldname = "Name"
				}

				details = append(details, tide.InfoDetails{
					strings.Replace(fieldname, " ", "", -1),
					strings.TrimSpace(value),
				})
			}
		}

		if validHeader {
			return
		}
	}

	return "other", nil, errors.New("not a valid theme or plugin file")
}
