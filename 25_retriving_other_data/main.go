package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		Method:        r.Method,
		URL:           r.URL,
		Submissions:   r.Form,
		Header:        r.Header,
		Host:          r.Host,
		ContentLength: r.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
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
