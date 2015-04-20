package sessions

import (
//	"fmt"
	"net/http"
	"html/template"


//	"golang.org/x/net/context"
//	"golang.org/x/oauth2"
//	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
//	"google.golang.org/appengine/user"
//	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
//	"google.golang.org/appengine/urlfetch"
//	"google.golang.org/cloud"
//	"google.golang.org/cloud/storage"

	"github.com/astaxie/beego/session"
)

type pageContent struct {
	Username string
	Data1 string
	Data2 string
	Data3 string
}

var (
	globalSessions *session.Manager	//global memory space for the session manager
	rootTmpl = template.Must(template.ParseFiles("templates/base.go.html", "templates/rootcontent.go.html"))
)

func init() {
	//TODO look at https://github.com/astaxie/beego/blob/master/session/sess_mem_test.go
	//to see if it is necessary to have this full configuration for this setup
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid",
														"enableSetCookie,omitempty": true,
														"gclifetime":180,
														"maxLifetime": 180,
														"secure": false,
														"sessionIDHashFunc": "sha1",
														"sessionIDHashKey": "",
														"cookieLifeTime": 180,
														"providerConfig": ""}`)	//since we are using memory, we don't need a provider config
	go globalSessions.GC()

	//Serve static files/content. ex. css and favicon
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("public/static"))))

	//Serve the pages
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/createsession", createSessionHandler)
	http.HandleFunc("/destroysession", destroySessionHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//open the session
	sess, _ := globalSessions.SessionStart(w,r)
//	defer sess.SessionRelease(w) //not implemented in beego?

	//check if the session is already created
	var pc pageContent
	uname := sess.Get("username")
	if uname != nil {
		pc.Username = sess.Get("username").(string)
		pc.Data1 = sess.Get("data1").(string)
		pc.Data2 = sess.Get("data2").(string)
		pc.Data3 = sess.Get("data3").(string)
	}

//	log.Infof(c, "Session ID: %v", sess.SessionID())
	rootTmpl.Execute(w, pc)
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {

	sess, _ := globalSessions.SessionStart(w,r)

	//Use the session Id as the user for now
	sess.Set("username", sess.SessionID())

	//Store the information the user set in the form
	sess.Set("data1", r.FormValue("data1"))
	sess.Set("data2", r.FormValue("data2"))
	sess.Set("data3", r.FormValue("data3"))

	//debug
	c := appengine.NewContext(r)
	log.Infof(c, "Session ID: %v", sess.SessionID())

	http.Redirect(w,r, "/", 302 )
}

func destroySessionHandler(w http.ResponseWriter, r *http.Request) {

	//get session info
	sess, _ := globalSessions.SessionStart(w, r)

	//destroy the session if it actually exists
	uname := sess.Get("username")
	if uname != nil {
		globalSessions.SessionDestroy(w,r)
	}

	http.Redirect(w,r, "/", 302 )
}

