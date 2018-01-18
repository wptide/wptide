package api

import (
	"net/url"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/xwp/go-tide/src/tide"
	"errors"
)

type Client struct {
	user *tide.Auth
}

// Authenticate creates a call to a Tide API instance and attempts to authenticate a user.
func (c *Client) Authenticate(clientId, clientSecret, authEndpoint string) error {

	// Create form data to send via the HTTP request.
	form := url.Values{
		"api_key":    {clientId},
		"api_secret": {clientSecret},
	}
	body := bytes.NewBufferString(form.Encode())

	// Send a POST request to Tide to authenticate.
	rsp, err := http.Post(authEndpoint, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// Read the response into a bytes buffer.
	body_byte, _ := ioutil.ReadAll(rsp.Body)

	// Attempt to unmarshal the JSON string into an Auth object.
	var auth *tide.Auth
	json.Unmarshal(body_byte, &auth)

	if rsp.StatusCode != http.StatusOK || auth == nil || auth.AccessToken == "" {
		return errors.New("Could not authenticate user.")
	}

	c.user = auth

	return nil
}

// SendPayload sends authenticated requests to a Tide API instance.
//
// `method` is POST or GET for Tide API.
// `endpoint` is the particular Tide endpoint to send the payload to.
// `token` is the Auth.AccessToken of an authenticated user object.
// `data` is a JSON encoded string to send with the payload (or empty).
func (c Client) SendPayload(method, endpoint, data string) (string, error) {
	var req *http.Request
	var err error

	req, err = http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(data)))

	if c.user != nil && c.user.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.user.AccessToken))
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	if ( resp.StatusCode < 200 || resp.StatusCode > 299 ) {
		return "", errors.New("Unexpected status code.")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), err
}
