package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type Cars struct {
	Id    int
	Value string
	Tag   string
}

func main() {

	cars := Cars{
		Id:    1,
		Value: "NISSAN",
		Tag:   `<script>alert("XSS should be happen");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", cars)
	if err != nil {
		log.Fatalln(err)
	}
}
