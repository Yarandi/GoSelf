package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/yarandi/", yarandi)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "you get the index page")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is me")
}

func yarandi(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Yarandi")
	if err != nil {
		log.Fatal("error in executing template", err)
	}
}
