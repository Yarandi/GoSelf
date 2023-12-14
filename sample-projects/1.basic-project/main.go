package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404", http.StatusNotFound)
	}
	if req.Method != "GET" {
		http.Error(res, "method not accepted", http.StatusNotFound)
	}
	fmt.Fprintf(res, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404", http.StatusNotFound)
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParsForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "address: %s\n", address)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting Server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
