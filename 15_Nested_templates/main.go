package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	pageTitle := "Nested Template"
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", pageTitle)
	if err != nil {
		log.Fatalln(err)
	}
}
