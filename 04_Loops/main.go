package main

import (
	"go/types"
	"log"
)

func main(){

	//slice range
	// animals := []string{"Dog","Horses","Cat","Giraffe"}
	// for i, animal := range animals {
	// 	log.Println(i,animal)
	// }

	// //map range
	// animals := make(map[string]string)

	// animals["Horse"] = "Charlie"
	// animals["Cat"] = "Flufle"
	// animals["Dog"] = "Max"

	// for i,v := range animals{
	// 	log.Println(i, v)
	// }

	type User struct {
		Name     string
		Family   string
		Email	 string
		Age      int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 32})
	users = append(users, User{"hamed", "Smith", "john@smith.com", 29})
	users = append(users, User{"Mary", "Smith", "john@smith.com", 59})
	users = append(users, User{"Jack", "Smith", "john@smith.com", 21})
	users = append(users, User{"John", "Smith", "john@smith.com", 36})

	for _, value := range users{
		log.Println(value.Name, value.Family, value.Email, value.Age)
	}


}