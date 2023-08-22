package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Set up the HTTP server to listen on port 8080
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

// Handler function for the "/" route
func foo(w http.ResponseWriter, req *http.Request) {
	// Get the request context
	ctx := req.Context()

	// Add some key-value pairs to the context to simulate user information
	ctx = context.WithValue(ctx, "user", "root")
	ctx = context.WithValue(ctx, "pass", "1234")
	ctx = context.WithValue(ctx, "port", 3606)

	// Call the getDbAccess function to simulate a database access operation
	res, err := getDbAccess(ctx)
	if err != nil {
		// If an error occurs, send an HTTP error response with the status code 408 (Request Timeout)
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}
	// Write the response content to the http.ResponseWriter
	fmt.Fprintln(w, res)
}

// Simulates a database access operation
func getDbAccess(ctx context.Context) (int, error) {
	// Create a new context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Create a channel to communicate with a goroutine
	ch := make(chan int)
	go func() {
		// Extract the port number from the context
		uid := ctx.Value("port").(int)
		// Simulate some work with a sleep
		time.Sleep(2 * time.Second)
		// Check if the context has expired due to timeout
		if ctx.Err() != nil {
			log.Fatalln("Timeout occurred!") // Print a log message
			return
		}
		// Send the port number to the channel
		ch <- uid
	}()

	// Use a select statement to wait for context expiration or channel value
	select {
	case <-ctx.Done():
		return 0, ctx.Err() // Return context error if context is done
	case i := <-ch:
		return i, nil // Return the port number
	}
}
