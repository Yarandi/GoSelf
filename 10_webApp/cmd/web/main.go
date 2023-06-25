package main

import (
	"fmt"
	"net/http"

	handler "github.com/Yarandi/go-course/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("starting Application on port %s ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
 