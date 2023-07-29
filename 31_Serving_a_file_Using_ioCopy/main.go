package main

import (
	"io"
	"net/http"
	"os"
)

// --------------------------------------------------------------------------------
// io.Copy allows us to copy from a reader to a writer.
// We can use io.Copy to read from a file and then write the file to the response.
// --------------------------------------------------------------------------------
func main() {

	//define routes
	http.HandleFunc("/", dog)
	http.HandleFunc("/tedd.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="tedd.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("tedd.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
