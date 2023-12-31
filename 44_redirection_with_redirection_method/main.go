package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("your request method at foo is: ", req.Method, "\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bar is: ", req.Method, "\n")
	//http.Redirect(w, req, "/", http.StatusSeeOther) // code = 303
	//http.Redirect(w, req, "/", http.StatusTemporaryRedirect) // code = 307
	http.Redirect(w, req, "/", http.StatusMovedPermanently) // code = 301
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Print("you request method at barred is: ", req.Method, "\n")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
