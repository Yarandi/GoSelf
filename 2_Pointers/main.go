package main

import "log"

func main() {

	var str string = "Green"	
	//original value
	log.Println("str value is", str)

	changeValueByPointer(&str)

	//log manipulled data
	log.Println("str value is", str)

}

func changeValueByPointer(s *string){

   newValue := "Red"

   *s = newValue

}