package main

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	//get the cookie
	c, err := r.Cookie("session")
	if err != nil {
		//create a cookie
		//SID = uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.New().String(),
		}
		http.SetCookie(w, c)
	}

	//if the user exists already, get user
	var u User
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = User{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
