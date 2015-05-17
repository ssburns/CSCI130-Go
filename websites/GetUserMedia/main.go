package main

import (
//	"fmt"
	"net/http"
)

func init() {

	//	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))


	http.HandleFunc("/", rootHandler)
}

func rootHandler( w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "public/index.html")
}