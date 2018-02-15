package phpcs

import (
	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"os/exec"
	"io/ioutil"
	"errors"
	"github.com/wptide/pkg/tide"
	"github.com/wptide/pkg/storage"
	"strings"
	"github.com/wptide/pkg/log"
	"github.com/wptide/pkg/audit/ingest"
	"fmt"
	"strconv"
	"syscall"
	"os"
)

var (
	// Using exec.Command as a variable so that we can mock it in tests.
	execCommand = exec.Command
)

type Processor struct {
	Standard       string            // The standard to run for this process.
	PostProcessors []audit.Processor // A post process to run on the results.
	Parallel       int               // Number of parallel phpcs cli processes (phpcs setting).
	Encoding       string
	RuntimeSet     []string // String array in format "key value" per entry.
	ReportFile     string
	ReportFilePath string
	FilesPath      string
}

func (p Processor) Kind() string {
	if p.standard() == "" {
		return "phpcs_undefined"
	}
	return "phpcs_" + p.standard()
}

func (p Processor) standard() string {
	return strings.TrimSpace(strings.ToLower(p.Standard))
}

// Process will run a phpcs audit against a code base and add to *results.
// Note: It checks *result first to see if the message has been ingested and if its processing a theme.
func (p *Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// Check if phpcs is supposed to run or exit.
	if ! audit.CanRunAudit(p, result) {
		return
	}

	log.Log(msg.Title, "Running PHPCS Audit...")

	// Get some critical information.
	ingestObj, ok := r["ingest"].(*ingest.Processor);
	if ! ok || ingestObj.Dest == "" {
		r[p.Kind()+"Error"] = errors.New("could not find path to code")
		r[p.Kind()] = nil

		log.Log(msg.Title, r[p.Kind()+"Error"])
		return
	}

	p.FilesPath = ingestObj.Dest + "/unzipped"

	var err error
	p.ReportFilePath, p.ReportFile, err = p.reportPath(r)
	if err != nil {
		r[p.Kind()+"Error"] = errors.New("could not determine report file path")
		r[p.Kind()] = nil

		log.Log(msg.Title, r[p.Kind()+"Error"])
		return
	}

	// Set some defaults if not defined.
	if p.Parallel < 1 {
		p.Parallel = 1
	}

	if p.Encoding == "" {
		p.Encoding = "utf-8"
	}

	cmdName := "phpcs"
	cmdArgs := []string{
		"--extensions=php",
		"--ignore=*/vendor/*,*/node_modules/*",
		"--standard=" + p.standard(),
		"--encoding=" + p.Encoding,
		"--basepath=.",
		"--report=json",
		"--report-json=" + p.ReportFilePath,
		"--parallel=" + strconv.Itoa(p.Parallel),
	}

	for _, pair := range p.RuntimeSet {
		split := strings.Split(pair, " ")
		if len(split) == 2 {
			cmdArgs = append(cmdArgs, "--runtime-set")
			cmdArgs = append(cmdArgs, split[0])
			cmdArgs = append(cmdArgs, split[1])
		}
	}

	cmdArgs = append(cmdArgs, p.FilesPath)
	cmdArgs = append(cmdArgs, "-q")

	// Prepare the command and set the stdOut pipe.
	cmd := execCommand(cmdName, cmdArgs...)

	cmdReader, err := cmd.StdoutPipe()
	p.errCheck(err, result)

	// Start the command.
	err = cmd.Start()
	p.errCheck(err, result)

	// Read stdout pipe.
	// Note: phpcs does not write to Stderr so nothing to read there.
	resultBytes, err := ioutil.ReadAll(cmdReader)
	p.errCheck(err, result)

	// Wait for command to exit and stdio to be read.
	r[p.Kind()+"ExitCode"] = 0
	exitErr := cmd.Wait()

	// Grab the Exit Code.
	if exitErr, ok := exitErr.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			r[p.Kind()+"ExitCode"] = status.ExitStatus()
		}
	}

	log.Log(msg.Title, fmt.Sprintf("phpcs output:\n %s", strings.TrimSpace(string(resultBytes))))

	// We already have a reference to the report file, so letss upload and get the storage reference in a result.
	log.Log(msg.Title, "Uploading "+p.Kind()+" results to remote storage.")
	fullResults, err := p.uploadToStorage(r)
	p.errCheck(err, result)

	// Initialise the result and set the "Full" entry to the uploaded file.
	r[p.Kind()] = &tide.AuditResult{}
	if fullResults != nil {
		auditResults := r[p.Kind()].(*tide.AuditResult)
		auditResults.Full = fullResults.Full
	}

	// Run post processes.
	fileReader, err := os.Open(p.ReportFilePath)
	p.errCheck(err, result)
	for _, proc := range p.PostProcessors {
		if proc, ok := proc.(audit.PostProcessor); ok && err == nil {
			fileReader.Seek(0, 0) // Rewind file.
			proc.Parent(p)
			proc.SetReport(fileReader)
			log.Log(msg.Title, "Post-Processing: "+proc.Kind())
			proc.Process(msg, &r)
		}
	}

	// If our Details are empty, then we assume we want Full here.
	emptyResult := tide.AuditResult{}
	if r[p.Kind()].(*tide.AuditResult).Details == emptyResult.Details {
		full := r[p.Kind()].(*tide.AuditResult).Full
		auditResults := r[p.Kind()].(*tide.AuditResult)
		auditResults.Details.Type = full.Type
		auditResults.Details.Key = full.Key
		auditResults.Details.BucketName = full.BucketName
	}

	log.Log(msg.Title, fmt.Sprintf("phpcs (%s) process completed with exit code: %d\n", p.standard(), r[p.Kind()+"ExitCode"]))
}

// @todo Replace with "OK-pattern interface"
func (p Processor) errCheck(err error, result *audit.Result) {
	r := *result
	if err != nil {
		r[p.Kind()+"Error"] = err
		log.Log("phpcs ("+p.standard()+") Error:", r[p.Kind()+"Error"])
	}
}

func (p *Processor) uploadToStorage(r audit.Result) (*tide.AuditResult, error) {

	var results *tide.AuditResult

	// Get required variables from results.
	storageProvider, providerOk := r["fileStore"].(*storage.StorageProvider)

	if ! providerOk {
		return nil, errors.New("could not get fileStore to upload to")
	}

	if p.ReportFilePath == "" || p.ReportFile == "" {
		var err error
		p.ReportFilePath, p.ReportFile, err = p.reportPath(r)

		if err != nil {
			r[p.Kind()+"Error"] = errors.New("could not determine report file path")
			r[p.Kind()] = nil

			log.Log(p.Kind(), r[p.Kind()+"Error"])
			return nil, r[p.Kind()+"Error"].(error)
		}
	}

	sP := *storageProvider
	err := sP.UploadFile(p.ReportFilePath, p.ReportFile)

	if err == nil {
		results = &tide.AuditResult{
			Full: struct {
				Type       string `json:"type,omitempty"`
				Key        string `json:"key,omitempty"`
				BucketName string `json:"bucket_name,omitempty"`
			}{
				Type:       sP.Kind(),
				Key:        p.ReportFile,
				BucketName: sP.CollectionRef(),
			},
		}
	}

	return results, err
}

func (p Processor) reportPath(r audit.Result) (string, string, error) {
	tempFolder, folderOk := r["tempFolder"].(string)
	checksum, checksumOk := r["checksum"].(string)

	if ! folderOk {
		return "", "", errors.New("no tempFolder to write results to before upload to fileStore")
	}

	if ! checksumOk {
		return "", "", errors.New("there was no checksum to be used for filenames")
	}

	storageRef := checksum + "-" + p.Kind() + "-full.json"
	filename := strings.TrimRight(tempFolder, "/") + "/" + storageRef

	return filename, storageRef, nil
}
