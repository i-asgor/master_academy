package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var google0authConfig = &oauth2.Config{
	ClientID:     "465885968547-n2su4qbg9e6t44nsq783be5i7j99et30.apps.googleusercontent.com",
	ClientSecret: "K0k7BKVB3q7j97unNkoouobd",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/cloud-platform"},
	RedirectURL: "http://localhost:8080/callback",
	Endpoint:    google.Endpoint,
}
var randomState = "random"

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := google0authConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := google0authConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	// fmt.Println(token)
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	// fmt.Println(resp.Body)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Response: %s", content)
	fmt.Println(string(content))

	// for i := 0; i < len(content); i++ {
	// 	fmt.Println(content[i]) // This logs uint8 and prints numbers
	// }

	// fmt.Println(reflect.TypeOf(content))
	// fmt.Println("done")
}
