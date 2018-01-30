package lighthouse

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"fmt"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/xwp/go-tide/src/tide"
	"github.com/xwp/go-tide/src/storage"
	"strings"
)

var (
	// Using exec.Command as a variable so that we can mock it in tests.
	execCommand = exec.Command
	writeFile   = ioutil.WriteFile
)

type Processor struct{}

func (p Processor) Kind() string {
	return "lighthouse"
}

// Process will run a lighthouse audit against a theme and to *results.
// Note: It checks *result first to see if the message has been ingested and if its processing a theme.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// Check if lighthouse is supposed to run or exit.
	if ! audit.CanRunAudit(p, result) {
		return
	}

	var results *tide.LighthouseSummary

	// Note: This assumes the shell script `lh` is in $PATH and contains the following command:
	// `lighthouse --quiet --chrome-flags="--headless --disable-gpu --no-sandbox" --output=json --output-path=stdout $@`
	cmdName := "lh"
	cmdArgs := []string{fmt.Sprintf("https://wp-themes.com/%s", msg.Slug)}

	// Prepare the command and set the stdOut pipe.
	cmd := execCommand(cmdName, cmdArgs...)

	cmdReader, err := cmd.StdoutPipe()
	errCheck(err, result)

	errReader, err := cmd.StderrPipe()
	errCheck(err, result)

	// Start the command.
	err = cmd.Start()
	errCheck(err, result)

	// Read stdout pipe.
	resultBytes, err := ioutil.ReadAll(cmdReader)
	errCheck(err, result)

	// Read stderr pipe.
	errorBytes, err := ioutil.ReadAll(errReader)
	errCheck(err, result)

	if len(errorBytes) > 0 {
		r["lighthouse"] = nil
		r["lighthouseError"] = errors.New("lighthouse command failed: " + string(errorBytes))
		return
	}

	// Wait for command to exit and stdio to be read.
	err = cmd.Wait()
	errCheck(err, result)

	// Unmarshal the body response into a LightHouseReport object.
	err = json.Unmarshal(resultBytes, &results)
	errCheck(err, result)

	auditResult := &tide.AuditResult {}

	// Upload and get full results.
	fullResults, err := uploadToStorage(r, resultBytes)
	errCheck(err, result)

	// @todo Decide if `Details` should only contain results from `reportCategories` and update structure in tide/item.go.
	if fullResults != nil {
		auditResult.Full = fullResults.Full
		auditResult.Details.Type = fullResults.Full.Type
		auditResult.Details.Key = fullResults.Full.Key
		auditResult.Details.BucketName = fullResults.Full.BucketName
	}

	auditResult.Summary = &tide.AuditSummary{
		LighthouseSummary: results,
	}

	r["lighthouse"] = auditResult
}

func errCheck(err error, result *audit.Result) {
	r := *result
	if err != nil {
		r["lighthouseError"] = err
	}
}

func uploadToStorage(r audit.Result, buffer []byte) (*tide.AuditResult, error) {

	var results *tide.AuditResult

	// Get required variables from results.
	storageProvider, providerOk := r["fileStore"].(*storage.StorageProvider)
	tempFolder, folderOk := r["tempFolder"].(string)
	checksum, checksumOk := r["checksum"].(string)

	if ! providerOk {
		return nil, errors.New("could not get fileStore to upload to")
	}

	if ! folderOk {
		return nil, errors.New("no tempFolder to write results to before upload to fileStore")
	}

	if ! checksumOk {
		return nil, errors.New("there was no checksum to be used for filenames")
	}

	storageRef := checksum + "-lighthouse-full.json"
	filename := strings.TrimRight(tempFolder, "/") + "/" + storageRef

	err := writeFile(filename, buffer, 0644)
	if err != nil {
		return nil, errors.New("could not write lighthouse audit to tempFolder")
	}

	sP := *storageProvider
	err = sP.UploadFile(filename, storageRef)

	if err == nil {
		results = &tide.AuditResult{
			Full: struct {
				Type       string `json:"type,omitempty"`
				Key        string `json:"key,omitempty"`
				BucketName string `json:"bucket_name,omitempty"`
			}{
				Type:       sP.Kind(),
				Key:        storageRef,
				BucketName: sP.CollectionRef(),
			},
		}
	}

	return results, err
}
