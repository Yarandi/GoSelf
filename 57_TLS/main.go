package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTPS!"))
	})

	// Specify paths to your TLS certificate and private key files
	certFile := "/etc/letsencrypt/live/example.com/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/example.com/privkey.pem"

	// Start the HTTPS server
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		//panic(err)
	}
}
