package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Hamed-key", "this is from Yarandi")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>code</h1>")
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
