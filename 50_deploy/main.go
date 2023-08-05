package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	//for {
	fmt.Fprintf(w, "Hello from deployed app on macos!!")
	//time.Sleep(time.Second * 5)
	//}
}
