package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error
var tpl *template.Template

type User struct {
	id    int
	name  string
	email string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
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
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/del", del)
	http.HandleFunc("/drop", drop)

	http.HandleFunc("/showTables", showTables)

	http.ListenAndServe(":8080", nil)
}

func showTables(w http.ResponseWriter, r *http.Request) {
	showTables := "SHOW Tables;"
	rows, err := db.Query(showTables)
	checkErr(err)

	//query
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		checkErr(err)
		//s += name + "\t" + email + "\n"
		fmt.Fprintln(w, tableName)
	}

	//err = tpl.ExecuteTemplate(w, "tables.gohtml", s)
	//checkErr(err)

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
	err = tpl.ExecuteTemplate(w, "index.gohtml", users)
	checkErr(err)
	//fmt.Fprintln(w, users)
}

func create(w http.ResponseWriter, r *http.Request) {
	createTableQuery := `CREATE TABLE test4 (name VARCHAR(20))`
	_, err = db.Exec(createTableQuery)
	checkErr(err)

	//fmt.Fprintln(w, "CREATED TABLE customer")
	http.Redirect(w, r, "/showTables", http.StatusSeeOther)
	return
}

func insert(w http.ResponseWriter, r *http.Request) {
	insertTableQuery := `Insert INTO test VALUES ("Hamed")`
	_, err = db.Exec(insertTableQuery)
	checkErr(err)
	fmt.Fprintln(w, "Value Inserted successfully")
}

func update(w http.ResponseWriter, r *http.Request) {
	updateTableQuery := `UPDATE test SET name="Yarandi" WHERE name="Hamed")`
	_, err = db.Exec(updateTableQuery)
	checkErr(err)
	fmt.Fprintln(w, "Value Updated successfully")

}

func del(w http.ResponseWriter, r *http.Request) {
	delQuery := `DELETE FROM TEST WHERE name="Hamed"`
	_, err = db.Exec(delQuery)
	checkErr(err)
	fmt.Fprintln(w, "Value deleted successfully")

}

func drop(w http.ResponseWriter, r *http.Request) {
	dropQuery := `DROP TABLE test`
	_, err = db.Exec(dropQuery)
	checkErr(err)
	fmt.Fprintln(w, "table dropped successfully")

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
