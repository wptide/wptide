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

	// Expected audits.
	var availableResults int
	expectedAudits, ok := r["audits"].([]string)

	if !ok {
		expectedAudits = []string{}
		r["audits"] = expectedAudits
	}

	// Only ingest if we do not know if this is a theme or we don't have a specific
	// route given. This will have to be calculated from checksum if not.
	isCollection := util.IsCollectionEndpoint(msg.ResponseAPIEndpoint)

	if ! isCollection {

		//audits := r["audits"].([]string)
		client := r["client"].(*tide.ClientInterface)
		response, _ := p.getResults(client, msg.ResponseAPIEndpoint)
		var results tide.Item
		err := json.Unmarshal([]byte(response), &results)
		if err != nil {
			r["ingest"], _ = json.Marshal(p)
			r["ingestError"] = err
			return
		}

		// Attach existing audit item to the result.
		r["tideItem"], _ = json.Marshal(results)

		// Attach checksum to the result.
		r["checksum"] = results.Checksum

		for _, a := range expectedAudits {
			if _, ok := results.Results[a]; ok {
				availableResults += 1
			}
		}

		// If we have results for all the audits there is nothing to do, unless its forced.
		if availableResults == len(expectedAudits) && ! msg.Force {
			r["ingest"], _ = json.Marshal(p)
			r["ingestError"] = errors.New("no new audits to run")
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
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(msg.SourceURL))

	tempFolder, ok := r["tempFolder"].(string)

	if !ok {
		r["ingest"] = nil
		r["ingestError"] = errors.New("can't extract without a `temptFolder`")
		return
	}

	p.Dest = tempFolder + "/lh-" + base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	err := srcMgr.PrepareFiles(p.Dest)
	if err != nil {
		r["ingestError"] = err
	}

	// Attach calculated checksum to the results.
	r["checksum"] = srcMgr.GetChecksum()
	r["files"] = srcMgr.GetFiles()

	r["ingest"], _ = json.Marshal(p)
}

func (p Processor) getResults(client *tide.ClientInterface, endpoint string) (string, error) {
	c := *client
	return c.SendPayload("GET", endpoint, "")
}
