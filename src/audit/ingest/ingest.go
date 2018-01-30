package ingest

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/util"
	"github.com/xwp/go-tide/src/source"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/xwp/go-tide/src/source/zip"
	"errors"
	"github.com/xwp/go-tide/src/tide"
	"strings"
	"github.com/xwp/go-tide/src/log"
)

var (
	srcMgr source.Source
)

type Processor struct {
	Dest  string   `json:"dest"`
	Files []string `json:"files"`
}

func (p Processor) Kind() string {
	return "ingest"
}

// Process takes *result and sends relevant results to the Tide API.
func (p *Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	log.Log(msg.Title, "Ingesting...")

	// Expected audits that need to be run. E.g. lighthouse.
	var availableResults int
	expectedAudits, ok := r["audits"].([]string)

	// If no audits are passed then set expectedAudits to nil.
	// It will flow still flow through remaining processes if there is other work to do.
	if !ok {
		expectedAudits = nil
		r["audits"] = expectedAudits
	}

	// Only ingest if we do not know if this is a theme or we don't have a specific
	// route given. This will have to be calculated from checksum if not.
	isCollection := util.IsCollectionEndpoint(msg.ResponseAPIEndpoint)

	// Get the Tide API client. We will use it for collections and specific routes.
	client := r["client"].(*tide.ClientInterface)

	// If its not a collection we need to attempt to get results from Tide API.
	if ! isCollection {

		// Fetch results from Tide API.
		response, _ := p.getResults(client, msg.ResponseAPIEndpoint)
		response = cleanResults(response)

		// Attempt to unmarshal the results.
		var results tide.Item
		err := json.Unmarshal([]byte(response), &results)
		if err != nil {
			r["ingest"], _ = json.Marshal(p)
			r["ingestError"] = err

			log.Log(msg.Title, r["ingestError"])
			return
		}

		// Attach existing audit item to the result.
		r["tideItem"] = &results

		// Attach checksum to the result.
		r["checksum"] = results.Checksum

		// See if results exist for the expectedAudits to run.
		for _, a := range expectedAudits {
			if _, ok := results.Results[a]; ok {
				availableResults += 1
			}
		}

		// If we have results for all the audits there is nothing to do, unless its forced.
		//if availableResults == len(expectedAudits) && ! msg.Force && results.CodeInfo.Type != "" {
		if availableResults == len(expectedAudits) && ! msg.Force {
			r["ingest"], _ = json.Marshal(p)
			r["ingestError"] = errors.New("no new audits to run")

			log.Log(msg.Title, r["ingestError"])
			return
		}
	}

	// If srcMgr is not set at package level, then set it.
	if srcMgr == nil {
		switch source.GetKind(msg.SourceURL) {
		case "zip":
			srcMgr = zip.NewZip(msg.SourceURL)
		}
	}

	// Without a Source manager we can't continue.
	if srcMgr == nil {
		r["ingest"] = nil
		r["ingestError"] = errors.New("could not handle source url")

		log.Log(msg.Title, r["ingestError"])
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(msg.SourceURL))

	tempFolder, ok := r["tempFolder"].(string)

	if !ok {
		r["ingest"] = nil
		r["ingestError"] = errors.New("can't extract without a `temptFolder`")

		log.Log(msg.Title, r["ingestError"])
		return
	}

	p.Dest = tempFolder + "/lh-" + base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	err := srcMgr.PrepareFiles(p.Dest)
	if err != nil {
		r["ingestError"] = err
		log.Log(msg.Title, r["ingestError"])
	}

	// Attach calculated checksum to the results.
	checksum := srcMgr.GetChecksum()
	r["checksum"] = checksum

	log.Log(msg.Title, "Project checksum: `" + checksum + "`")

	// Attach file paths.
	r["files"] = srcMgr.GetFiles()

	// If collection route, now that we have the checksum its worth checking the API for existing results.
	if isCollection && checksum != "" {
		// Fetch results from Tide API.
		response, _ := p.getResults(client, strings.TrimRight(msg.ResponseAPIEndpoint, "/")+"/"+checksum)
		response = cleanResults(response)

		// Attempt to unmarshal the results.
		var results tide.Item
		err := json.Unmarshal([]byte(response), &results)
		if err == nil {
			r["tideItem"] = &results
		}
	}

	r["ingest"] = p
}

func (p Processor) getResults(client *tide.ClientInterface, endpoint string) (string, error) {
	c := *client
	return c.SendPayload("GET", endpoint, "")
}

func cleanResults(in string) string {
	in = strings.Replace(in, `,"code_info":""`, "", -1)
	in = strings.Replace(in, `,"results":[]`, "", -1)
	return in
}
