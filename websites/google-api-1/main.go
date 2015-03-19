package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
//	"golang.org/x/oauth2"
//	"google.golang.org/appengine/urlfetch"
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

type ClientSecret2 struct {
	Web struct {
		Auth_uri string `json: "auth_uri"`
		Client_secret string `json: "client_secret"`
		Token_uri string `json: "token_uri"`
		Client_email string `json: "client_email"`
		Redirect_uris []string `json: "redirect_uris"`
		Client_x509_cert_url string `json: "client_x509_cert_url"`
		Client_id string `json: "client_id"`
		Auth_provider_x509_cert_url string `json: "auth_provider_x509_cert_url"`
		Javascript_origins []string `json: "javascript_origins"`
	}`json: "web"`
}

func main() {
	//First lets see what we need to read the json file
	cSecret := make(map[string]map[string]interface{},0)

	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		fmt.Println("File error:", err)
		os.Exit(1)
	}


	fmt.Println("*******************************")

	var clientSecret ClientSecret
	err = json.Unmarshal(file, &clientSecret)
	fmt.Println(clientSecret.Web.Client_secret)
	fmt.Println(err)


	fmt.Println("*******************************")

	json.Unmarshal(file, &cSecret)

	for key, value := range(cSecret["web"]) {
		fmt.Println(key, ":", value)
	}
}
