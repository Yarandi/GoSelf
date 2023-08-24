package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) MarshalJSON() ([]byte, error) {
	fullName := p.FirstName + " " + p.LastName
	data := map[string]interface{}{
		"full_name": fullName,
		"age":       p.Age,
	}
	return json.Marshal(data)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	if fullName, ok := jsonData["full_name"].(string); ok {
		names := splitFullName(fullName)
		p.FirstName = names[0]
		p.LastName = names[1]
	}

	if age, ok := jsonData["age"].(float64); ok {
		p.Age = int(age)
	}

	return nil
}

func splitFullName(fullName string) []string {
	return []string{"", ""} // You can implement this function based on your needs
}

func main() {
	p1 := Person{FirstName: "Alice", LastName: "Johnson", Age: 30}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Custom Marshaled JSON:")
	fmt.Println(string(jsonData))

	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Unmarshaled Person:")
	fmt.Println(p2)
}
