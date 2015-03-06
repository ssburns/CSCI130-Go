/*
HTML/TEMPLATE
The html/template package is part of the Go standard library.
We can use html/template to keep the HTML in a separate file,
allowing us to change the layout of our edit page
without modifying the underlying Go code.
*/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
	Fname string
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	strfname := "You"
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, Fname: strfname}, nil
}

/*
BEFORE EXTRACTING renderTemplate:
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}
*/

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// template.ParseFiles
// reads the contents of an html file and returns a *template.Template
// t.Execute
// executes the template, writing the generated HTML to the http.ResponseWriter

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

// defaultHandler ADD BY TODD MCLEOD
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	//directory path
	dir, _ := os.Open(".")
	defer dir.Close()

	fileInfos, _ := dir.Readdir(-1)

	body := ""

	for _, fi := range fileInfos {
		name := fi.Name()
		name = name[0:len(name)-4]
		body += "<a href='http://localhost:8080/view/" + name +"'>"+"http://localhost:8080/view/" + name +"</a>" + "\n"

		//fmt.Println("-",fi.Name())
	}

	fmt.Fprintf(w, body)

//	fmt.Fprintf(w, "<a href='http://localhost:8080/view/testpage'>"+
//				"http://localhost:8080/view/testpage</a>")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
