package main

import (
//	"fmt"
	"net/http"

)

func init() {

	http.HandleFunc("/", rootHandler)
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "basic_video_call.html")
}