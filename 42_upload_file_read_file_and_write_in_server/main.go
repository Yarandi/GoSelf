package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	var s string
	//check if method is post
	if r.Method == http.MethodPost {

		fmt.Println(r.Method)
		//open the file
		f, h, err := r.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//for learning
		fmt.Printf("file: %t, header: %d, error: %t", f, h, err)

		//read the file
		b, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(b)

		//check if directory exists or not (is not necessary check)
		// Check if the directory already exists
		_, err = os.Stat("home/documents/user/upload")
		if err != nil {
			fmt.Println("Directory already exists.")
			return
		}
		//make directory
		err = os.MkdirAll("home/documents/user/upload", 0750)
		if err != nil && !os.IsExist(err) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//store on server
		dst, err := os.Create(filepath.Join("home", "documents", "user", "upload", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//defer dst.Close()

		_, err = dst.Write(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	//set header and render the template
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)

}
