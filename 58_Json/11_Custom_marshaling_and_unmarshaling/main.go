package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Custom JSON marshaling logic
func (p Person) MarshalJSON() ([]byte, error) {
	type CustomPerson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	customPerson := CustomPerson{
		Name: p.FirstName + " " + p.LastName,
		Age:  p.Age,
	}

	return json.Marshal(customPerson)
}

// Custom JSON unmarshaling logic
func (p *Person) UnmarshalJSON(data []byte) error {
	type CustomPerson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var customPerson CustomPerson
	err := json.Unmarshal(data, &customPerson)
	if err != nil {
		return err
	}

	// Split the name into first name and last name
	names := splitName(customPerson.Name)
	p.FirstName = names[0]
	p.LastName = names[1]
	p.Age = customPerson.Age

	return nil
}

func splitName(name string) []string {
	return []string{"John", "Doe"} // Dummy implementation, you can use strings.Split() here
}

func main() {
	person := Person{
		FirstName: "Alice",
		LastName:  "Johnson",
		Age:       30,
	}

	// Marshal the Person struct
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("Custom Marshaled JSON:", string(jsonData))

	// Unmarshal the JSON back to a Person struct
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Custom Unmarshaled Person:", decodedPerson)
}
