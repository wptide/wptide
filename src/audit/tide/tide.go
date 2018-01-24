package tide

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"github.com/pkg/errors"
	"github.com/xwp/go-tide/src/util"
	"encoding/json"
)

type Processor struct{}

func (p Processor) Kind() string {
	return "tide"
}

// Process takes *result and sends relevant results to the Tide API.
func (p Processor) Process(msg message.Message, result *audit.Result) {

	// Cannot perform indexing on the *result directly, so assigning pointer to a local variable.
	r := *result

	var client *tide.ClientInterface

	// Attempt to get the client.
	if r["client"] != nil {
		client = r["client"].(*tide.ClientInterface)
	}

	// If we have no client there is nothing we can do.
	if client == nil {
		r["tideError"] = errors.New("Tide API client not provided")
		return
	}

	if r["audits"] == nil {
		r["tideError"] = errors.New("there are no audits to send results for")
		return
	}

	// Get the right endpoint to send results to.
	var endpoint string
	if util.IsCollectionEndpoint(msg.ResponseAPIEndpoint) && r["tideItem"] != nil {
		// If we found an item on a collection endpoint we will use the checksum endpoint.
		endpoint = msg.ResponseAPIEndpoint + "/" + r["checksum"].(string)
	} else {
		// Use the given endpoint.
		endpoint = msg.ResponseAPIEndpoint
	}

	audits := r["audits"].([]string)
	resultMap := map[string]tide.AuditResult{}

	// Loop through the audits and assign results to the result map.
	for _, audit := range audits {
		if r[audit] != nil {
			if a, ok := r[audit].(*tide.AuditResult); ok {
				resultMap[audit] = *a
			}
		}
	}

	// Prepare Item for Payload.
	payloadItem := &tide.Item{
		Title:       msg.Title,
		Version:     "",
		Checksum:    "",
		Visibility:  msg.Visibility,
		ProjectType: msg.ProjectType,
		SourceUrl:   msg.SourceURL,
		SourceType:  msg.SourceType,
		CodeInfo:    tide.CodeInfo{},
		Results:     resultMap,
	}

	if checksum, ok := r["checksum"].(string); ok {
		payloadItem.Checksum = checksum
	}


	payload, _ := json.Marshal(payloadItem)

	_, err := p.sendResults(client, endpoint, string(payload))
	if err != nil {
		r["tideError"] = err
	}
}

func (p Processor) sendResults(client *tide.ClientInterface, endpoint, data string) (string, error) {
	c := *client
	return c.SendPayload("POST", endpoint, data)
}
