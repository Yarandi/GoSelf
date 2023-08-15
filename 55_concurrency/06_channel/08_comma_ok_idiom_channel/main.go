package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		c <- 37
		c <- 38
		c <- 39
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok) //37 true
	v, ok = <-c
	fmt.Println(v, ok) //38 true
	v, ok = <-c
	fmt.Println(v, ok) //39 true

}
