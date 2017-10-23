package repo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
	"strconv"
	"errors"
)

// Info struct represents the "info" key in the API response.
// Key names are camel-case thought they appear lowercase in the API.
// json.Unmarshal will deal with this appropriately.
type Info struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Results int `json:"results"`
}

// Plugin struct represents the plugin objects within the plugins array in the API response.
type Plugin struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	DownloadLink string `json:"download_link"`
	Version      string `json:"version"`
	LastUpdated  string `json:"last_updated"`
}

// PluginsResponse struct represents the response from the API.
type PluginsResponse struct {
	Info    Info
	Plugins []Plugin
}

func RequestPlugins(browse string, perPage int, page int) (PluginsResponse, error) {
	requestUrl := "https://api.wordpress.org/plugins/info/1.1/"

	// formValues is an array of query parameters passed to the requestUrl.
	formValues := []string{
		"action=query_plugins",
		"request[browse]=" + browse, // Can be `popular`, `featured`, `updated`, `new`
		"request[per_page]=" + strconv.Itoa(perPage),
		"request[page]=" + strconv.Itoa(page),
		"request[fields][sections]=0",
		"request[fields][description]=0",
		"request[fields][short_description]=0",
		"request[fields][tested]=0",
		"request[fields][requires]=0",
		"request[fields][rating]=0",
		"request[fields][ratings]=0",
		"request[fields][downloaded]=0",
		"request[fields][last_updated]=1",
		"request[fields][homepage]=0",
		"request[fields][tags]=0",
		"request[fields][donate_link]=0",
		"request[fields][contributors]=0",
		"request[fields][compatibility]=0",
		"request[fields][versions]=0",
		"request[fields][version]=0",
		"request[fields][screenshots]=0",
		"request[fields][stable_tag]=0",
		"request[fields][download_link]=1",
	};

	requestParams := bytes.NewBufferString(strings.Join(formValues, "&"))

	response, err := http.Post(requestUrl, "application/x-www-form-urlencoded", requestParams)
	if err != nil {
		return PluginsResponse{}, errors.New("could not retrieve plugins from WordPres.org API")
	}

	defer response.Body.Close()
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return PluginsResponse{}, errors.New("could not read response body")
	}

	results := PluginsResponse{}
	json.Unmarshal([]byte(bodyByte), &results)

	return results, nil
}
