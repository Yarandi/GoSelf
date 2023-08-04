package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Session",
		Value: "Some Value",
	})

	fmt.Fprintln(w, `<h1><a href="/read">read cookie</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Session")
	if err != nil {
		http.Redirect(w, r, "set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>your cookie is </br> %v </h1> <h1><a href="/expire">expire</h1>`, cookie)
}

func expire(w http.ResponseWriter, r *http.Request) {

	//get the cookie
	c, err := r.Cookie("Session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //delete the cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
