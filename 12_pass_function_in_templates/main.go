package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//tpl = template.Must(template.ParseFiles("index.gohtml"))
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	m := map[int]string{
		1: "hamed",
		2: "saeed",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", m)
	if err != nil {
		log.Fatal(err)
	}
}
