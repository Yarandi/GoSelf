package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!!!!")
}
