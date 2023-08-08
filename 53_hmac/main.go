package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

//-----------------------------------------------------------------------------
//// Create an HMAC instance using SHA-256 hash function Use for authentication
//-----------------------------------------------------------------------------

func main() {
	var str string = "hamedYar"
	result := getcode(str)
	fmt.Println(result)
}

func getcode(str string) string {
	// Create an HMAC instance using SHA-256 hash function and secret key "salt"
	hash := hmac.New(sha256.New, []byte("salty"))

	// Write the string s into the HMAC instance
	io.WriteString(hash, str)

	// Calculate the HMAC hash
	return fmt.Sprintf("%x", hash.Sum(nil))
}
