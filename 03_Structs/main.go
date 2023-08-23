package main

import "log"

type User struct{
	FirstName string
	Age int	
}

func (m *User) getFirstName() string {
	return m.FirstName
}

func main(){

	var myVar User
	myVar.FirstName = "hamed"

	myVar2 := User {
		FirstName: "Mina",
		Age: 32,
	}

	log.Println(myVar.getFirstName())
	log.Println(myVar2.getFirstName())

}