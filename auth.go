package quizongo

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type authError struct {
	Err map[string]string
}

type AccessToken struct {
	Access_token string
	Token_type string
	Expires_in int
	Scope string
	User_id string
}

var auth_url = "https://quizlet.com/oauth/token"

func AuthURL(client_id, redirect_uri string, options map[string]string) (output string) {
	auth_url, _ := url.Parse("https://quizlet.com/authorize")
	auth_query := auth_url.Query()
	auth_query.Add("client_id", client_id)
	auth_query.Add("redirect_uri", redirect_uri)
	auth_query.Add("response_type", "code") //Based on quizlet documentation.

	for key, value := range options {
		auth_query.Add(key, value)
	}

	auth_url.RawQuery = auth_query.Encode()

	return auth_url.String()
}

func ObtainAccessToken(client_id, client_secret, code, redirect_uri string) (output *AccessToken, error error) {
	client := &http.Client{Transport: getTransport()}

	data := url.Values{"code": {code}, "redirect_uri": {redirect_uri}, "grant_type": {"authorization_code"}}
	request, error := http.NewRequest("POST", auth_url, strings.NewReader(data.Encode()))

	if error != nil {
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(client_id, client_secret)

	// make the request.
	response, error := client.Do(request)

	if error != nil {
		return
	}

	//check whether the response is a bad request
	if response.StatusCode != 200 {
		collection := new(authError)
		error = parseResponse(response, collection)
		error = errors.New(collection.Err["error"] + ": " + collection.Err["error_description"])
	} else {
		collection := new(AccessToken)
		error = parseResponse(response, collection)
		output = collection
	}
	return
}
