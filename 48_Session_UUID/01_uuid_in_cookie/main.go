package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {

	//get the cookie
	c, err := req.Cookie("Session")
	if err != nil {
		//create a new cookie
		nc := &http.Cookie{
			Name:     "Session",
			Value:    uuid.New().String(),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, nc)
	}
	fmt.Println(c)
}
