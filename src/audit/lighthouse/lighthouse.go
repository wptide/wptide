package lighthouse

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"fmt"
	"os/exec"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
)

// Report contains a truncated version of the full report.
type Report struct {
	ReportCategories  []ReportCategory `json:"reportCategories"`
}

// ReportCategory contains the results for each category without the audits.
type ReportCategory struct {
	Name        string  `json:"name"`
	Weight      float32 `json:"weight"`
	Description string  `json:"description"`
	Id          string  `json:"id"`
	Score       float32 `json:"score"`
}

type Processor struct{}

// Process will run a lighthouse audit against a theme and to *results.
// Note: It checks *result first to see if the message has been ingested and if its processing a theme.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// Note: This assumes the shell script `lh` is in $PATH and contains the following command:
	// `lighthouse --quiet --chrome-flags="--headless --disable-gpu --no-sandbox" --output=json --output-path=stdout $@`
	cmdName := "lh"
	cmdArgs := []string{fmt.Sprintf("https://wp-themes.com/%s", msg.Slug)}

	// Prepare the command and set the stdOut pipe.
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		r["lighthouse"] = err
		return
	}

	// Start the command.
	err = cmd.Start()
	if err != nil {
		log.Println(os.Stderr, "Error starting Cmd", err)
		r["lighthouse"] = err
		return
	}

	// Read stdout pipe.
	resultBytes, _ := ioutil.ReadAll(cmdReader)

	// Wait for command to exit and stdio to be read.
	err = cmd.Wait()
	if err != nil {
		log.Println(os.Stderr, "Error waiting for Cmd", err)
		r["lighthouse"] = err
		return
	}

	// Unmarshal the body response into a LightHouseReport object.
	var results *Report
	err = json.Unmarshal(resultBytes, &results)
	if err != nil {
		log.Println(os.Stderr, "Error validating report", err)
		r["lighthouse"] = err
		return
	}

	r["lighthouse"] = results
}
