package phpcs

import (
	"io"
	"io/ioutil"
	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
	"encoding/json"
	"github.com/xwp/go-tide/src/phpcompat"
	"strings"
	"errors"
	"github.com/wptide/pkg/storage"
	"sort"
)

var (
	writeFile   = ioutil.WriteFile
)

// PhpcsSummary implements audit.PostProcessor.
type PhpCompat struct {
	Report        io.Reader
	ParentProcess audit.Processor
	resultPath    string
	resultFile    string
}

func (p PhpCompat) Kind() string {
	return "phpcs_phpcompat"
}

func (p *PhpCompat) Process(msg message.Message, result *audit.Result) {
	r := *result

	// Byte buffer for reading from the report file.
	var byteSummary []byte
	var err error

	if p.Report != nil {
		// Read the report into the byte buffer.
		byteSummary, _ = ioutil.ReadAll(p.Report)
	}

	// We need an empty result to compare nil values.
	emptyResult := tide.AuditResult{}

	// Get current audit results from parent.
	auditResults, ok := r[p.ParentProcess.Kind()].(*tide.AuditResult)

	// If we can't get the results there is nothing to do.
	if ! ok {
		errMsg := "could not get results from parent process"
		r[p.Kind()+"Error"] = errors.New(errMsg)
		return
	}

	// Full "raw" results from the PHPCS process.
	var fullResults tide.PhpcsResults
	err = json.Unmarshal(byteSummary, &fullResults)
	if err != nil {
		errMsg := "could not get phpcs results"
		r[p.Kind()+"Error"] = errors.New(errMsg)
		auditResults.Error += "\n" + errMsg
		return
	}

	brokenVersions := []string{}
	sources := make(map[string]map[string]interface{})

	// Iterate files and only get summary data.
	for filename, data := range fullResults.Files {
		for _, sniffMessage := range data.Messages {

			if _, ok := sources[sniffMessage.Source]; !ok {
				sources[sniffMessage.Source] = make(map[string]interface{})
				sources[sniffMessage.Source]["files"] = make(map[string]interface{})
				broken := phpcompat.BreaksVersions(sniffMessage.Source)
				brokenVersions = mergeVersions(brokenVersions, broken)
				sources[sniffMessage.Source]["breaks"] = broken
			}
			sources[sniffMessage.Source]["files"].(map[string]interface{})[filename] = sniffMessage
		}
	}

	// Will Marshall without error because the sources map gets initialized with `make()`
	res, _ := json.Marshal(sources)

	p.resultFile, p.resultPath, err = p.summaryPath(r, "detail")

	err = writeFile(p.resultPath, res, 0644)
	if err != nil {
		errMsg := "could not write PHPCompatibility details to disk"
		r[p.Kind()+"Error"] = errors.New(errMsg)
		auditResults.Error += "\n" + errMsg
		return
	}

	details, err := p.uploadResults(r)
	if err != nil {
		errMsg := "could not write PHPCompatibility details to file store"
		r[p.Kind()+"Error"] = errors.New(errMsg)
		auditResults.Error += "\n" + errMsg
		return
	}

	// We don't want to override, so only add a summary if there is no summary.
	if auditResults.Details == emptyResult.Details {
		auditResults.Details = details.Details
	}

	auditResults.CompatibleVersions = getCompatibleVersions(brokenVersions)
}

func (p *PhpCompat) SetReport(report io.Reader) {
	p.Report = report
}

func (p *PhpCompat) Parent(parent audit.Processor) {
	p.ParentProcess = parent
}

func (p PhpCompat) uploadResults(r audit.Result) (results *tide.AuditResult, err error) {

	var storageProvider *storage.StorageProvider
	var ok bool

	if p.resultFile == "" || p.resultPath == "" {
		err = errors.New("no result path given")
		return
	}

	if storageProvider, ok = r["fileStore"].(*storage.StorageProvider); !ok {
		err = errors.New("could not get fileStore to upload to")
		return
	}

	sP := *storageProvider
	err = sP.UploadFile(p.resultPath, p.resultFile)

	if err == nil {
		results = &tide.AuditResult{
			Details: struct {
				Type       string `json:"type,omitempty"`
				Key        string `json:"key,omitempty"`
				BucketName string `json:"bucket_name,omitempty"`
				*tide.PhpcsResults
				*tide.LighthouseResults
			}{
				Type:       sP.Kind(),
				Key:        p.resultFile,
				BucketName: sP.CollectionRef(),
			},
		}
	}

	return
}

func (p PhpCompat) summaryPath(r audit.Result, fileSuffix string) (filename, path string, err error) {

	var checksum, tempFolder string
	var ok bool

	if tempFolder, ok = r["tempFolder"].(string); ! ok {
		err = errors.New("no tempFolder to write results to before upload to fileStore")
		return
	}

	if checksum, ok = r["checksum"].(string); ! ok {
		err = errors.New("there was no checksum to be used for filenames")
		return
	}

	filename = checksum + "-" + p.Kind()
	if fileSuffix != "" {
		filename += "-" + fileSuffix + ".json"
	} else {
		filename += ".json"
	}

	path = strings.TrimRight(tempFolder, "/") + "/" + filename

	return
}

func mergeVersions(n ...[]string) []string {
	merged := []string{}

	for _, slice := range n {
		for _, value := range slice {
			if ! contains(merged, value) {
				merged = append(merged, value)
			}
		}
	}

	return merged
}

// getCompatibleVersions removes the broken versions from the PhpMajorVersions.
func getCompatibleVersions(broken []string) []string {

	compatible := []string{}

	for _, version := range phpcompat.PhpMajorVersions() {
		if ! contains(broken, version) && ! contains(compatible, version) {
			compatible = append(compatible, version)
		}
	}

	sort.Strings(compatible)
	return compatible
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
