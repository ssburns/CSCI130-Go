package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	addr      = flag.String("addr", ":8080", "http service address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("github.com/gary.burd.info/go-websocket-chat", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func main() {
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "home.html")))

	//create the hub
	h := newHub()	//because chan have to be made with make
	go h.run()

	//Register Address with the Mux

	//This is where people go to connect to the chat
	http.HandleFunc("/", homeHandler)	//HandleFunc -> homeHandler has the arguments for (ResponseWriter, *Request)

	//This is where the chat client connects
	http.Handle("/ws", wsHandler{h: h})	//Handle -> wsHandler has the method ServeHTTP(ResponseWriter, *Request) (the Handler interface)

	//Serve!
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
