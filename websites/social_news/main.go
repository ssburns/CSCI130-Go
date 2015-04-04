package hello

import (
	"fmt"
	"bytes"
	"net/http"
	"io"
	"time"
//	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"

//	"appengine"
//	"appengine/user"
//	"appengine/datastore"
)

type webSubmission struct {
	Title string
	Link string
	SubmitBy string
	Thread int64
	SubmitDateTime time.Time
	SubmissionDesc string
}

type webUser struct {
	Uuid uint64
	Nickname string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/create2", create2Handler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/update2", updateHandler2)
	http.HandleFunc("/delete", deleteHandler)

}



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	http.ServeFile(w,r, "public/templates/create.html")
}

func create2Handler(w http.ResponseWriter, r *http.Request) {
	strTitle := r.FormValue("title")
	strLink := r.FormValue("link")
//	fmt.Fprint(w, "go title: ", strTitle)
//	fmt.Fprint(w, "link: ", strLink)

	//TODO error checking for the title and link

	c := appengine.NewContext(r)
	u := user.Current(c)

	//TODO error checking for user?

	newSubmission := webSubmission {
		Title: strTitle,
		Link: strLink,
		SubmitBy: u.String(),
		Thread: 123,	//TODO create random thread id
	}


	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "webSubmission", nil), &newSubmission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
//	http.Redirect(w, r, "/", http.StatusFound)
	fmt.Fprintf(w, "Thank you for your submission")
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("webSubmission").
		Filter("SubmitBy =", "test@example.com")

	b := new(bytes.Buffer)
	for t := q.Run(c); ; {
		var x webSubmission
		key, err := t.Next(&x)
		if err == datastore.Done{
			break
		}
		if err != nil{
			serveError(c,w,err)
			fmt.Fprintf(w, "nope %v", err.Error())
			return
		}
		fmt.Fprintf(b,"Key=%v\nwebSubmission=%#v\n\n", key, x)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.Copy(w,b)
//	fmt.Fprint(w, "Hello, read!")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	http.ServeFile(w,r, "public/templates/update.html")
}

func updateHandler2(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("webSubmission").
		//		Filter("SubmitBy =", "test@example.com")
		Filter("Title =", "test1")

	//TODO cleanup checks whatever for emtpy or otherwise other than 1 result
	t := q.Run(c)

	var x webSubmission
	key, err := t.Next(&x)
	if err != nil{
		serveError(c,w,err)
		fmt.Fprintf(w, "nope %v", err.Error())
		return
	}

	//Update field
	x.Link = r.FormValue("link")


	//Store back
	_, err2 := datastore.Put(c, key, &x)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

//	fmt.Fprint(w, "Hello Update")
	http.Redirect(w,r,"/read",http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("webSubmission").
		//		Filter("SubmitBy =", "test@example.com")
		Filter("Title =", "test1")

	//TODO cleanup checks whatever for emtpy or otherwise other than 1 result
	t := q.Run(c)

	var x webSubmission
	key, err := t.Next(&x)
	if err != nil{
		serveError(c,w,err)
		fmt.Fprintf(w, "nope %v", err.Error())
		return
	}

	err2 := datastore.Delete(c, key)
	if err != nil {
		serveError(c,w,err2)
		return
	}

	http.Redirect(w,r,"/read", http.StatusFound)
}

func serveError(c context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
//	c.Errorf("%v", err)
}
