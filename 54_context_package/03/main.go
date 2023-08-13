package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "user", "root")
	ctx = context.WithValue(ctx, "pass", "1234")
	ctx = context.WithValue(ctx, "port", 3606)

	res, err := getDbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout) //408
	}
	fmt.Fprintln(w, res)
}

func getDbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		uid := ctx.Value("port").(int)
		time.Sleep(4 * time.Second)
		if ctx.Err() != nil {
			log.Fatalln("Timeout occurred!") // Add a log message
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}
