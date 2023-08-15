package main

import "fmt"

func main() {
	//two way channel
	c := make(chan int)

	//Create a channel that can only be used for receiving
	rc := make(<-chan int)

	//Create a channel that can only be used for sending
	sc := make(chan<- int)

	fmt.Println("-------------------")
	fmt.Printf("c \t%T\n", c)
	fmt.Printf("rc \t%T\n", rc)
	fmt.Printf("sc \t%T\n", sc)

	//general to specific converts
	fmt.Println("-------------------")
	fmt.Printf("c \t%T\n", (<-chan int)(c))
	fmt.Printf("c \t%T\n", (chan<- int)(c))

}
