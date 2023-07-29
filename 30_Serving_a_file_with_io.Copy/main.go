package main

import (
	"io"
	"net/http"
)

// --------------------------------------------------------------------------------
// io.Copy allows us to copy from a reader to a writer.
// We can use io.Copy to read from a file and then write the file to the response.
// --------------------------------------------------------------------------------
func main() {

	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<!--not serving from our server-->
    <div><h1><em>not serving from our server</em></h1></div>
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
