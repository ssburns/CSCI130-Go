package main

import (
//	"fmt"
	"net/http"
	"html/template"
	"time"
)

var (
	tmpl = template.Must(template.ParseFiles("root.html"))
)


func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view", viewHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("username")
	var username string

	if cookie != nil {
		username = cookie.Value
	}

	tmpl.ExecuteTemplate(w, "root.html", username)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	//Create cookie
	expiration := time.Now().Add(3 * time.Minute)
	cookie := http.Cookie{Name: "username", Value:name, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w,r,"/", 302)
}
