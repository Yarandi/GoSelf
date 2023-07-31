package main

import (
	"io"
	"net/http"
)

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
	http.ServeFile(w, r, "tedd.jpg")
}
