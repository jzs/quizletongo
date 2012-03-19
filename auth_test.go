package quizongo

import (
	"testing"
)

func TestAuthURL(t *testing.T) {
}

func TestObtainAccessTokenValid(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/oauth/token", dummyAccessTokenResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	auth_url = dummy_server.URL + "/oauth/token"

	token_output, _ := ObtainAccessToken("abc", "secret", "def", "www.my_app.com")

	if token_output.Access_token != "46a54395f3d1108feca56c7f6ca8dd3d" {
		t.Error("Access token does not match.")
	}
	if token_output.Scope != "read" {
		t.Error("Scope does not match.")
	}
	if token_output.User_id != "USERNAME" {
		t.Error("Username does not match.")
	}
	if token_output.Expires_in != 3600 {
		t.Error("Expiry date wrong")
	}
}

func TestObtainAccessTokeninvalid(t *testing.T) {
}

//test data
var dummyAccessTokenResponse string = `
{
	"access_token": "46a54395f3d1108feca56c7f6ca8dd3d",
	"token_type": "bearer",
	"expires_in": 3600,
	"scope": "read",
	"user_id": "USERNAME"
} `

var dummyAccesTokenErrorResponse string = `
{
	"error": "invalid_request",
	"error_description": "Invalid grant_type parameter or parameter missing" 
} `
