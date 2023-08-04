package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//check if cookie exists or not
	cookie, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	//converting value from string to int
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
	count++

	//converting value from int into string again
	cookie.Value = strconv.Itoa(count)

	//set again the cookie
	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}
