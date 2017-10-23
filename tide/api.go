package api

import (
	"os"
	"net/http"
	"net/url"
	"bytes"
	"strconv"
)

var client *http.Client

func AddRevision(title, sourceUrl string) error {

	// @todo Remove this after testing.
	return nil

	endpoint := os.Getenv("TIDE_API_BASE")
	// api_key := os.Getenv("TIDE_API_KEY")
	// api_secret := os.Getenv("TIDE_API_SECRET")

	// Get the HTTP client.
	if client == nil {
		client = &http.Client{}
	}

	formValues := url.Values{
		"title":         {title},
		"content":       {title},
		"revision_meta": {"{\"audit_tag\":\"syncserver\",\"audit_project\":\"SyncServer\"}"},
		"visibility":    {"public"},
		"source_url":    {sourceUrl},
		"source_type":   {"zip"},
	}

	// Create a new request.
	request, err := http.NewRequest("POST", endpoint+"tide/v2/audit_revision/", bytes.NewBufferString(formValues.Encode()))
	if err != nil {
		return err
	}

	// Add JWT Bearer Token
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE0OTAzMzc0OTAsImlzcyI6Imh0dHA6XC9cL3RpZGUuZGV2IiwiZXhwIjoxNDkyOTI5NDkwLCJkYXRhIjp7ImNsaWVudCI6eyJpZCI6MiwidHlwZSI6IndwX3VzZXIifX19.jdWFQsEufHSGFTxbC85Nm6v1Tn6Kf3du2HSygiY-pUo"
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	// Setup form headers
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(formValues.Encode())))

	// Execute the request
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
