package main

import (
//	"fmt"
	"net/http"
	"html/template"
)

var (
	tmpl = template.Must(template.ParseFiles("root.html", "view.html"))
)


func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view", viewHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "root.html", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pageData := struct{Data1, Data2 string}{r.FormValue("name"), r.FormValue("link")}
	tmpl.ExecuteTemplate(w,"view.html", pageData)
}
