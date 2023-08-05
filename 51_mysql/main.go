package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

type User struct {
	id    int
	name  string
	email string
}

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/driving_school")
	checkErr(err)
	defer db.Close()

	checkErr(db.Ping())

	//routes
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "mysql is successfully connected!!!")
	checkErr(err)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,name,email FROM users")
	checkErr(err)
	defer rows.Close()

	var name, email string
	var id int
	var users []User
	//query
	for rows.Next() {
		err = rows.Scan(&id, &name, &email)
		checkErr(err)
		//s += name + "\t" + email + "\n"
		users = append(users, User{
			id:    id,
			name:  name,
			email: email,
		})
	}
	fmt.Fprintln(w, users)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
