package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	http.Handle("/faviocn.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "Go look at your terminal")
}