package quizongo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var host string = "https://api.quizlet.com"
var transport http.RoundTripper

var UseSSL bool = true

type Session struct {
	apikey string
}

func NewSession(apikey string) *Session {
	return &Session{apikey: apikey}
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

	if response.StatusCode >= 400 && response.StatusCode <= 500 {
		error = errors.New("Bad Request")
		return
	}

	//parse JSON
	error = json.Unmarshal(bytes, result)

	if error != nil {
		print(error.Error())
	}

	return
}

func (session Session) get(section string, params map[string]string, collection interface{}) (error error) {
	if params == nil {
		params = map[string]string{}
	}
	params["client_id"] = session.apikey
	return get(section, params, collection)
}

func (session Session) put(section string, params map[string]string, collection interface{}) (error error) {
	if params == nil {
		params = map[string]string{}
	}
	params["client_id"] = session.apikey
	return put(section, params, collection)
}

func (session Session) post(section string, params map[string]string, collection interface{}) (error error) {
	if params == nil {
		params = map[string]string{}
	}
	params["client_id"] = session.apikey
	return post(section, params, collection)
}

func (session Session) delete(section string, params map[string]string) (error error) {
	if params == nil {
		params = map[string]string{}
	}
	params["client_id"] = session.apikey
	return delete(section, params)
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

func put(section string, params map[string]string, collection interface{}) (error error) {
	client := &http.Client{Transport: getTransport()}

	request, error := http.NewRequest("PUT", setupEndpoint(section, params).String(), nil)
	response, error := client.Do(request)

	if error != nil {
		return
	}

	error = parseResponse(response, collection)

	return
}

func post(section string, params map[string]string, collection interface{}) (error error) {
	//client := &http.Client{Transport: getTransport()}

	//response, error := client.PostForm(setupEndpoint(section, params).String())
	return
}

func delete(section string, params map[string]string) (error error) {
	client := &http.Client{Transport: getTransport()}

	request, error := http.NewRequest("DELETE", setupEndpoint(section, params).String(), nil)
	response, error := client.Do(request)

	if error != nil {
		return
	}

	if response.StatusCode != 204 {
		error = errors.New("Delete failed")
		return
	}

	return
}
