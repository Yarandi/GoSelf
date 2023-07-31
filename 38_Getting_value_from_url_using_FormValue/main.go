package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favion.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	v := r.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}



