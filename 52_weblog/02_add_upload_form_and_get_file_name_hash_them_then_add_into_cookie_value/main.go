package main

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {

	c, err := r.Cookie("session")
	if err != nil {
		//create the cookie
		SID := uuid.NewString()
		c = &http.Cookie{
			Name:  "session",
			Value: SID,
		}
		//set the cookie
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {

	//values
	p1 := "rivers.jpg"
	p2 := "trees.jpg"
	p3 := "sands.jpg"

	//append name of the pictures in the session value, if it does not contain the name
	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c

}
