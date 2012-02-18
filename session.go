package quizongo

import (
	"net/http"
	"net/url"
	"errors"
	"io/ioutil"
	"encoding/json"
)

var host string = "https://api.quizlet.com"
var transport http.RoundTripper

var UseSSL bool = true

type Session struct {
	Apikey string
}

func NewSession(apikey string) *Session {
	return &Session{Apikey: apikey}
}

func getTransport() http.RoundTripper {
	if transport != nil {
		return transport
	}
	return http.DefaultTransport
}

func SetTransport(t http.RoundTripper) {
	transport = t
}

// construct the endpoint URL
func setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)

	if UseSSL {
		base_url.Scheme = "https"
	} else {
		base_url.Scheme = "http"
	}

	endpoint, _ := base_url.Parse("/2.0/" + path)

	query := endpoint.Query()
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
}

// parse the response
func parseResponse(response *http.Response, result interface{}) (error error) {
	// close the body when done reading
	defer response.Body.Close()

	//read the response
	bytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return
	}

	//parse JSON
	error = json.Unmarshal(bytes, result)

	if error != nil {
		print(error.Error())
	}

	//check whether the response is a bad request
	if response.StatusCode == 400 {
		error = errors.New("Bad Request")
	}

	return
}

func (session Session) get(section string, params map[string]string, collection interface{}) (error error) {
	if params == nil {
		params = map[string]string{}
	}
	params["client_id"] = session.Apikey
	return get(section, params, collection)
}

func get(section string, params map[string]string, collection interface{}) (error error) {
	client := &http.Client{Transport: getTransport()}

	response, error := client.Get(setupEndpoint(section, params).String())

	if error != nil {
		return
	}

	error = parseResponse(response, collection)

	return
}
