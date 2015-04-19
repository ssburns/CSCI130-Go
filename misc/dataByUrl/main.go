package main

import (
//	"fmt"
	"net/http"
	"html/template"
)

var (
	tmpl = template.Must(template.ParseFiles("root.html"))
)


func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "root.html", nil)
}

func
