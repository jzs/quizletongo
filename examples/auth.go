package main

import (
	"bitbucket.org/jzs/quizletongo"
	"os/exec"
	"fmt"
	"net/http"
	"log"
	"net"
)

var CLIENT_ID string = ""
var REDIRECT_URI string = "http://localhost:7000"
var SECRET string = ""
var lis net.Listener

func main() {
	// Get Auth url with scope read and CSRF state set.
	auth_url := quizongo.AuthURL(CLIENT_ID, REDIRECT_URI, map[string]string{"scope":"read", "state": "bobcat"})

	// Open browser for authorization. 
	// NOTE: This is only for OSX. On linux you should launch a specific
	// browser. E.g. Firefox or chromium etc.
	cmd := exec.Command("open", auth_url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Create tcp listener.
	lis, e := net.Listen("tcp", ":7000")
	if e != nil {
		panic(e)
	}

	// Handle response and close connection.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		values := r.URL.Query()

		if values.Get("error") != "" {
			fmt.Fprintf(w, "Error: %v\nDescription: %v", values.Get("error"), values.Get("error_description"))
			defer lis.Close()
			return
		}

		fmt.Fprintf(w, "You logged in successfully. You can now close your browser.")

		code := values.Get("code")

		token, err := quizongo.ObtainAccessToken(CLIENT_ID, SECRET, code, REDIRECT_URI)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Token: %v\nExpiring: %v\nUser Id: %v", token.Access_token, token.Expires_in, token.User_id)
		defer lis.Close()
	})

	// Start serving http.
	http.Serve(lis, nil)
}
