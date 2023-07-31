package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	q := r.FormValue("q")
	c := r.FormValue("c")
	w.Header().Set("Content-Type", "text-html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
		<label>q</label><input type="text" name="q"></br>
        <label>c</label><input type="text" name="c"></br>
		<input type="submit" value="Submit">
		</form></br> `+q+c)

}
