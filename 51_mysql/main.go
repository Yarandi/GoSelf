package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/driving_school")
	defer db.Close()
	checkErr(err)

	checkErr(db.Ping())

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "mysql is successfully connected!!!")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
