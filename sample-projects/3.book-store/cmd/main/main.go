package main

import (
	"3.book-store/pkg/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:9010", r)
}
