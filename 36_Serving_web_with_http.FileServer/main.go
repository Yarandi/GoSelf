package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	_, err := os.Open("tedd.jpg")
	if err != nil {
		http.Error(w, "not found", 404)
	}
}
