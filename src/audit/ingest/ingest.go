package ingest

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/util"
	"github.com/xwp/go-tide/src/source"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Processor struct{
	Dest string `json:"dest"`
	Files []string `json:"files"`
}

func (p Processor) Kind() string {
	return "ingest"
}

// Process takes *result and sends relevant results to the Tide API.
func (p *Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	// Only ingest if we do not know if this is a theme or we don't have a specific
	// route given. This will have to be calculated from checksum if not.
	isCollection := util.IsCollectionEndpoint(msg.ResponseAPIEndpoint)

	if ! isCollection {
		return
	}

	srcMgr := source.NewSourceManager(msg.SourceURL)

	hasher := sha256.New()
	hasher.Write([]byte(msg.SourceURL))
	p.Dest = "/tmp/lh-" + base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	err := srcMgr.PrepareFiles(p.Dest)
	if err != nil {
		r["ingestError"] = err
	}

	r["ingest"], _ = json.Marshal(p)
}
