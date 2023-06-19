package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson :=
		`[
		{
		  "first_name" : "Clark",
		  "last_name" : "Joy",
		  "hair_color" : "Black",
		  "has_dog" : true

		},
		{
		  "first_name" : "Betty",
		  "last_name" : "Fllor",
		  "hair_color" : "Brown",
		  "has_dog" : false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshaling json", err)
	}
	log.Printf("Unmarshalled %v", unmarshalled)

	//write a json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "hamed"
	m1.LastName = "yarandi"
	m1.HairColor = "brown"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Mina"
	m2.LastName = "Fam"
	m2.HairColor = "brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshaling", err)
	}
	fmt.Println(string(newJson))

}
