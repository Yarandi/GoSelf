package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "my-cookie",
		Value:  "the value is value",
		Domain: "",
	})
	fmt.Fprint(w, "Cookie written - check your browser")
	fmt.Fprintln(w, "Where should i check in firefox??!!")
}

func read(w http.ResponseWriter, req *http.Request) {

	cook, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //404
		return
	}
	fmt.Fprintln(w, "Your Cookie:", cook)
}
