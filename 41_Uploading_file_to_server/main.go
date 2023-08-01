package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		//open the file
		f, h, err := r.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//do not forget close file
		defer f.Close()

		//for you learning
		fmt.Println("\n File:", f, "\nHeader:", h, "\nError :", err)

		//now we need read file
		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//fmt.Fprintf(w, string(bs))

		err = tpl.ExecuteTemplate(w, "index.gohtml", string(bs))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
