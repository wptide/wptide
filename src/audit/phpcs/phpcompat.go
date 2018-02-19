package phpcs

import (
	"io"
	"io/ioutil"
	"github.com/wptide/pkg/audit"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/tide"
	"encoding/json"
	"fmt"
)

// PhpcsSummary implements audit.PostProcessor.
type PhpCompat struct {
	Report        io.Reader
	ParentProcess audit.Processor
}

func (p PhpCompat) Kind() string {
	return "phpcs_phpcompat"
}

func (p *PhpCompat) Process(msg message.Message, result *audit.Result) {
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

		sources := make(map[string]map[string]interface{})

		// Iterate files and only get summary data.
		for filename, data := range fullResults.Files {
			for _, msg := range data.Messages {
				if _, ok := sources[msg.Source]; !ok {
					sources[msg.Source] = make(map[string]interface{})
				}
				sources[msg.Source][filename] = msg
			}
		}

		res, _ := json.Marshal(sources)
		fmt.Println(string(res))

		summary := getPhpcsSummary(fullResults)

		auditResults.Summary = &tide.AuditSummary{
			PhpcsSummary: summary,
		}
	}
}


func (p *PhpCompat) SetReport(report io.Reader) {
	p.Report = report
}

func (p *PhpCompat) Parent(parent audit.Processor) {
	p.ParentProcess = parent
}
