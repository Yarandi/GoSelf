package main

import (
	"io"
	"net/http"
	"os"
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
	file, err := os.Open("tedd.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()
	//io.Copy(w, file)

	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)

}
