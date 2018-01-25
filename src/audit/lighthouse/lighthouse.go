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
)

var (
	// Using exec.Command as a variable so that we can mock it in tests.
	execCommand = exec.Command
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

	// @todo Write `resultBytes` to S3.

	// @todo Add full report after completing S3 component.
	auditResult := &tide.AuditResult{}
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
