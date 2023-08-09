package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "dbName", "mysql_database")
	ctx = context.WithValue(ctx, "userName", "products")
	ctx = context.WithValue(ctx, "password", "root")
	ctx = context.WithValue(ctx, "port", 6306)

	result := dbAccess(ctx)
	fmt.Fprintln(w, result)
}

func dbAccess(ctx context.Context) int {
	uname, ok := ctx.Value("port").(int)
	if !ok {
		log.Fatalln("there is something wrong in type assertion")
	}
	return uname
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
