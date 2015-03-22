package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"html/template"
//	"os"
	"log"
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	gaeLog "google.golang.org/appengine/log"
//	gmail "google.golang.org/api/gmail/v1"

)

type ClientSecret struct {
	Web WebType `json: "web"`
}

type WebType struct{
	Auth_uri string `json: "auth_uri"`
	Client_secret string `json: "client_secret"`
	Token_uri string `json: "token_uri"`
	Client_email string `json: "client_email"`
	Redirect_uris []string `json: "redirect_uris"`
	Client_x509_cert_url string `json: "client_x509_cert_url"`
	Client_id string `json: "client_id"`
	Auth_provider_x509_cert_url string `json: "auth_provider_x509_cert_url"`
	Javascript_origins []string `json: "javascript_origins"`
}

var conf = new(oauth2.Config)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/g_start", g_start)
	http.HandleFunc("/oauth2callback", oauth2callback)
	http.HandleFunc("/formResult", formHandler)
}

func oauth2callback( w http.ResponseWriter, r *http.Request) {


	//TODO: validate FormValue("state")

	code :=  r.FormValue("code")

	c := appengine.NewContext(r)
//	gaeLog.Infof(c, "State Val: %s", r.FormValue("state"))

	tok, err := conf.Exchange(c, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")


	fmt.Fprint(w, "No Autographs!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func g_start(w http.ResponseWriter, r *http.Request) {
	//Read the client_secret.json and parse the file so we can get our Google authorization
	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		fmt.Println("client_secret.json:", err)
	}

	var clientSecret ClientSecret
	err = json.Unmarshal(file, &clientSecret)
	if err != nil {
		fmt.Println("client_secret.json unmarshall err:", )
	}

	//TODO: update with ConfigFromJson
	conf = &oauth2.Config{
		ClientID: clientSecret.Web.Client_id,
		ClientSecret: clientSecret.Web.Client_secret,
		RedirectURL: clientSecret.Web.Redirect_uris[0],
		Scopes: []string {
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	if err != nil {
		log.Fatal(err)
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	http.Redirect(w, r, url, http.StatusFound)
//	fmt.Fprintf(w, "Visit the URL for the auth dialog: %v", url)
	// The following client will be authorized by the App Engine
	// app's service account for the provided scopes.
//	client := http.Client{Transport: opts.NewTransport()}
//	client.Get("...")
}

//func g_start(w http.ResponseWriter, r *http.Request) {
//
//	//Read the client_secret.json and parse the file so we can get our Google authorization
//	file, err := ioutil.ReadFile("./config/client_secret.json")
//	if err != nil {
//		fmt.Println("client_secret.json:", err)
//	}
//
////	var clientSecret ClientSecret
////	err = json.Unmarshal(file, &clientSecret)
////	if err != nil {
////		fmt.Println("client_secret.json unmarshall err:", )
////	}
////
////	conf := &oauth2.Config{
////		ClientID: clientSecret.Web.Client_id,
////		ClientSecret: clientSecret.Web.Client_secret,
////		RedirectURL: clientSecret.Web.Redirect_uris[0],
////		Scopes: []string {
////			"https://www.googleapis.com/auth/gmail.readonly",
////		},
////		Endpoint: google.Endpoint,
////	}
//
//	google.ConfigFromJSON()
//
//	//Redirect user to Google's consent page
//	url := conf.AuthCodeURL("state")
//	fmt.Fprintf(w, "Visit the URL for the auth dialog: %v", url)
//
//	//Handle the exchange code to initiate a transport
//	tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
//	if err != nil {
//		log.Fatalf("oauth2 exchange", err)
//	}
//
//	client := conf.Client(oauth2.NoContext, tok)
//	client.Get("...")
//
//	fmt.Fprint(w, "No Autographs!")
//
////	rootForm, err := ioutil.ReadFile("templates/prompt.html");
////	if err != nil {
////		http.NotFound(w, r)
////		return
////	}
////	fmt.Fprint(w, string(rootForm))
//}

var resultFile, _ = ioutil.ReadFile("templates/result.html");
var resultHtmlTemplate = template.Must(template.New("result").Parse(string(resultFile)))

func formHandler(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")

	var err error

	if strEntered == "Shawn" {
		err = resultHtmlTemplate.Execute(w, "You spelled it correctly")
	} else {
		err = resultHtmlTemplate.Execute(w, "...you spelled it wrong")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func main() {
//
//	file, err := ioutil.ReadFile("./config/client_secret.json")
//	if err != nil {
//		fmt.Println("File error:", err)
//		os.Exit(1)
//	}
//
//
//	fmt.Println("*******************************")
//
//	var clientSecret ClientSecret
//	err = json.Unmarshal(file, &clientSecret)
//	fmt.Println(clientSecret.Web.Client_secret)
//
//}
