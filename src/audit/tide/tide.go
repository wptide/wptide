package tide

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/tide"
	"github.com/pkg/errors"
	"github.com/xwp/go-tide/src/util"
	"encoding/json"
	"fmt"
	"strings"
	"reflect"
	"github.com/xwp/go-tide/src/env"
	"github.com/xwp/go-tide/src/log"
)

type PayloadMessage struct {
	tide.Item
	RequestClient string   `json:"request_client"`
	Project       []string `json:"project"`
}

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

	log.Log(msg.Title, "Preparing Tide API payload...")

	// If we have no client there is nothing we can do.
	if client == nil {
		r["tideError"] = errors.New("Tide API client not provided")
		log.Log(msg.Title, r["tideError"])
		return
	}

	if r["audits"] == nil {
		r["tideError"] = errors.New("there are no audits to send results for")
		log.Log(msg.Title, r["tideError"])
		return
	}

	var payloadItem *tide.Item
	var resultMap map[string]tide.AuditResult

	// If we have an item, lets use that.
	if item, ok := r["tideItem"].(*tide.Item); ok {
		payloadItem = item
		resultMap = item.Results
		log.Log(msg.Title, "Using existing audit for payload.")
	}

	// Get the right endpoint to send results to.
	var endpoint string
	var isCollection = util.IsCollectionEndpoint(msg.ResponseAPIEndpoint)
	if isCollection && payloadItem != nil {
		// If we found an item on a collection endpoint we will use the checksum endpoint.
		endpoint = strings.TrimRight(msg.ResponseAPIEndpoint, "/") + "/" + r["checksum"].(string)
	} else {
		// Use the given endpoint.
		endpoint = msg.ResponseAPIEndpoint
	}

	// If we don't have results, then make the map.
	if resultMap == nil {
		resultMap = make(map[string]tide.AuditResult)
	}

	audits := r["audits"].([]string)
	errorSlice := []string{}

	// Loop through the audits and assign results to the result map.
	for _, audit := range audits {
		if r[audit] != nil {
			if a, ok := r[audit].(*tide.AuditResult); ok && r[audit+"Error"] == nil {
				resultMap[audit] = *a
			} else {
				resultMap[audit] = tide.AuditResult{
					Error: "Error occurred during audit.",
				}
				errorSlice = append(errorSlice, audit)
			}
		} else {
			errorSlice = append(errorSlice, audit)
		}
	}

	if len(errorSlice) > 0 {
		r["auditErrors"] = errors.New(fmt.Sprintf("the following audits could not run: %s", strings.Join(errorSlice, ",")))
		log.Log(msg.Title, r["auditErrors"])
	}

	if len(resultMap) == 0 {
		r["tideError"] = errors.New("no results to send to API")
		r["tide"] = nil
		log.Log(msg.Title, r["tideError"])
		return
	}

	codeInfo, ok := r["info"].(*tide.CodeInfo)
	if !ok {
		codeInfo = &tide.CodeInfo{}
	}

	// If we don't have an existing item, create it.
	if payloadItem == nil {
		payloadItem = &tide.Item{
			Title:       msg.Title,
			Version:     "",
			Checksum:    "",
			Visibility:  msg.Visibility,
			ProjectType: msg.ProjectType,
			SourceUrl:   msg.SourceURL,
			SourceType:  msg.SourceType,
			CodeInfo:    *codeInfo,
		}
	}

	// Prepare Item for Payload.
	payloadItem.Results = resultMap

	if checksum, ok := r["checksum"].(string); ok {
		payloadItem.Checksum = checksum
	}

	// Validate the item and fill in missing fields.
	validItem, _ := getValidItem(isCollection, *payloadItem, msg, *result, *codeInfo)

	log.Log(msg.Title, "Sending payload to Tide API...")
	payload, _ := json.Marshal(validItem)
	_, err := p.sendResults(client, endpoint, string(payload))
	if err != nil {
		r["tideError"] = err
		log.Log(msg.Title, r["tideError"])
	}
}

func (p Processor) sendResults(client *tide.ClientInterface, endpoint, data string) (string, error) {
	c := *client
	return c.SendPayload("POST", endpoint, data)
}

func getValidItem(isCollection bool, item tide.Item, msg message.Message, result audit.Result, codeInfo tide.CodeInfo) (*PayloadMessage, error) {

	validItem := &PayloadMessage{}

	// Get possible existing item from results.
	resultItem, ok := result["tideItem"].(*tide.Item);
	if ! ok {
		resultItem = &tide.Item{}
	}

	resultSimpleInfo := tide.SimplifyCodeDetails(resultItem.CodeInfo.Details)
	simpleInfo := tide.SimplifyCodeDetails(codeInfo.Details)

	// Check CodeInfo
	validItem.CodeInfo, ok = fallbackValue(resultItem.CodeInfo, codeInfo).(tide.CodeInfo)

	// Some values can come straight from the provided item.
	validItem.Results = item.Results

	// If it can't be determined with code info, then get it from the message first. Allows changes.
	validItem.Visibility, ok = fallbackValue(msg.Visibility, resultItem.Visibility).(string)

	// Set taxonomy for this project.
	project, ok := fallbackValue(msg.Slug, simpleInfo.TextDomain).(string)
	if project != "" {
		validItem.Project = []string{project}
	}

	// Some values should only be set if it was a collection endpoint.
	if isCollection || msg.Force {
		validItem.Title, ok = fallbackValue(simpleInfo.Name, msg.Title).(string)
		validItem.Description, ok = fallbackValue(simpleInfo.Description, msg.Content).(string)
		validItem.ProjectType, ok = fallbackValue(resultItem.CodeInfo.Type, codeInfo.Type, msg.ProjectType).(string)
		validItem.Version, ok = fallbackValue(resultSimpleInfo.Version, simpleInfo.Version).(string)

		validItem.Checksum = item.Checksum
		validItem.SourceType = item.SourceType
		validItem.SourceUrl = item.SourceUrl

		// Set the request client on the API.
		defaultRequestClient := env.GetEnv("LH_DEFAULT_REQUEST_CLIENT=audit_server", "audit_server")
		validItem.RequestClient, ok = fallbackValue(msg.RequestClient, defaultRequestClient).(string)
	}

	return validItem, nil
}

func fallbackValue(value ...interface{}) interface{} {

	lastVal := reflect.TypeOf(value[0]).String()

	for _, val := range value {

		valType := reflect.TypeOf(val).String()
		if lastVal != valType {
			return nil
		}
	}

	for _, val := range value {

		valType := reflect.TypeOf(val).String()

		switch valType {

		case "tide.CodeInfo":
			if val.(tide.CodeInfo).Type != "" {
				return val.(tide.CodeInfo)
			}

		case "int":
			fallthrough
		case "int32":
			fallthrough
		case "int64":
			if val.(int) != 0 {
				return val.(int)
			}

		case "string":
			if val.(string) != "" {
				return val.(string)
			}

		default:
			return val
		}
	}
	return nil
}
