package main

import "fmt"

//unbuffered channel

// When a sender sends a value on an unbuffered channel, it blocks until a receiver receives that value.
func main() {

	ch := make(chan int)

	go func() {
		ch <- 36
	}()

	fmt.Println(<-ch)
}
