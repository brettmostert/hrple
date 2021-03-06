package auth

// func authenticate() {

// }

// package main

// import (
// 	"context"
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"time"

// 	"github.com/fatih/color"
// 	"github.com/skratchdot/open-golang/open"
// 	"golang.org/x/oauth2"
// )

// var (
// 	conf *oauth2.Config
// 	ctx  context.Context
// )

// func callbackHandler(w http.ResponseWriter, r *http.Request) {
// 	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

// 	// Use the authorization code that is pushed to the redirect
// 	// URL.
// 	code := queryParts["code"][0]
// 	log.Printf("code: %s\n", code)

// 	// Exchange will do the handshake to retrieve the initial access token.
// 	tok, err := conf.Exchange(ctx, code)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Token: %s", tok)
// 	// The HTTP Client returned by conf.Client will refresh the token as necessary.
// 	client := conf.Client(ctx, tok)

// 	resp, err := client.Get("http://some-server.example.com/")
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println(color.CyanString("Authentication successful"))
// 	}
// 	defer resp.Body.Close()

// 	// show succes page
// 	msg := "<p><strong>Success!</strong></p>"
// 	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
// 	fmt.Fprintf(w, msg)
// }

// func main() {
// 	ctx = context.Background()
// 	conf = &oauth2.Config{
// 		ClientID:     "cli",
// 		ClientSecret: "",
// 		Scopes:       []string{"openid", "profile"},
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  "https://oauth.example.com/dex/auth",
// 			TokenURL: "https://oauth.example.com/dex/token",
// 		},
// 		// my own callback URL
// 		RedirectURL: "http://127.0.0.1:9999/oauth/callback",
// 	}

// 	// add transport for self-signed certificate to context
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	sslcli := &http.Client{Transport: tr}
// 	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

// 	// Redirect user to consent page to ask for permission
// 	// for the scopes specified above.
// 	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

// 	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
// 	time.Sleep(1 * time.Second)
// 	open.Run(url)
// 	time.Sleep(1 * time.Second)
// 	log.Printf("Authentication URL: %s\n", url)

// 	http.HandleFunc("/oauth/callback", callbackHandler)
// 	log.Fatal(http.ListenAndServe(":9999", nil))

// }
