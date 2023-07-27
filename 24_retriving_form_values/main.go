package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------------
	//r.Form | get data from url (query string) and also get the body form data
	//-------------------------------------------------------------------------
	//-------------------------------------------------------------------------
	//r.PostForm | only get the body form data
	//-------------------------------------------------------------------------
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
