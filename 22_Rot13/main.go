package main

import (
	"fmt"
)

func rot13(s string) string {
	result := ""
	for _, char := range s {
		switch {
		case char >= 'a' && char <= 'z':
			// Rotate lowercase letters
			char = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			// Rotate uppercase letters
			char = 'A' + (char-'A'+13)%26
		}
		result += string(char)
	}
	return result
}

func main() {
	// Original text
	message := "A"

	// Apply ROT13 to the message
	encoded := rot13(message)

	// Decode the encoded message using ROT13 again
	decoded := rot13(encoded)

	fmt.Println("Original Message:", message)
	fmt.Println("Encoded Message:", encoded)
	fmt.Println("Decoded Message:", decoded)
}
