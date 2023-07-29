package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<!--not serving from our server-->
    <div><h1><em>not serving from our server</em></h1></div>
	<img src="https://ted.jpg">
	`)
}
