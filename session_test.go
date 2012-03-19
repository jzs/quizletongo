package quizongo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"fmt"
)

func closeDummyServer(dummy_server *httptest.Server) {
	transport = nil
	dummy_server.Close()
}

func createDummyServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	dummy_server := httptest.NewServer(http.HandlerFunc(handler))

	//change the host to use the test server
	SetTransport(&http.Transport{Proxy: func(*http.Request) (*url.URL, error) { return url.Parse(dummy_server.URL) }})

	//turn off SSL
	UseSSL = false

	return dummy_server
}

func returnDummyResponseForPath(path string, dummy_response string, t *testing.T) *httptest.Server {
	//serve dummy responses
	dummy_data := []byte(dummy_response)

	return createDummyServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
		fmt.Printf("path: %v", r.URL.Path)
			t.Error("Path doesn't match")
		}
		w.Write(dummy_data)
	})
}
