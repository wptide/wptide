package phpcs

import (
	"io"
	"io/ioutil"
	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
	"encoding/json"
)

// PhpcsSummary implements audit.PostProcessor.
type PhpcsSummary struct {
	Report        io.Reader
	ParentProcess audit.Processor
}

func (p PhpcsSummary) Kind() string {
	return "phpcs_summary"
}

func (p *PhpcsSummary) Process(msg message.Message, result *audit.Result) {
	r := *result

	var byteSummary []byte
	var err error

	if p.Report != nil {
		byteSummary, err = ioutil.ReadAll(p.Report)
		if err != nil {
			return
		}
	}

	emptyResult := tide.AuditResult{}
	auditResults, ok := r[p.ParentProcess.Kind()].(*tide.AuditResult)

	// If we can't get the results there is nothing to do.
	if ! ok {
		return
	}

	// We don't want to override, so only add a summary if there is no summary.
	if auditResults.Summary == emptyResult.Summary {

		var fullResults tide.PhpcsResults
		err := json.Unmarshal(byteSummary, &fullResults)
		if err != nil {
			auditResults.Error += "\n" + "could not get phpcs results"
			return
		}

		summary := &tide.PhpcsSummary{
			ErrorsCount:   fullResults.Totals.Errors,
			WarningsCount: fullResults.Totals.Warnings,
			FilesCount:    len(fullResults.Files),
			Files: make(map[string]struct {
				Errors   int `json:"errors"`
				Warnings int `json:"warnings"`
			}),
		}

		// Iterate files and only get summary data.
		for filename, data := range fullResults.Files {
			summary.Files[filename] = struct {
				Errors   int `json:"errors"`
				Warnings int `json:"warnings"`
			} {
				data.Errors,
				data.Warnings,
			}
		}

		auditResults.Summary = &tide.AuditSummary{
			PhpcsSummary: summary,
		}
	}
}

func (p *PhpcsSummary) SetReport(report io.Reader) {
	p.Report = report
}

func (p *PhpcsSummary) Parent(parent audit.Processor) {
	p.ParentProcess = parent
}
