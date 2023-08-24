package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25, Email: "bob@example.com"}
	p3 := Person{Name: "Charlie", Age: 22, Address: "123 Main St"}

	people := []Person{p1, p2, p3}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

//without omitEmpty tag result would be
//[{"name":"Alice","age":30,"email":"","address":""},{"name":"Bob","age":25,"email":"bob@example.com","address":""},{"name":"Charlie","age":22,"email":"","address":"123 Main St"}]

//with omitEmpty tag result would be
//[{"name":"Alice","age":30},{"name":"Bob","age":25,"email":"bob@example.com"},{"name":"Charlie","age":22,"address":"123 Main St"}]
