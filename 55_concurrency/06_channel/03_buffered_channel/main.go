package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 36
	ch <- 37

	fmt.Println(<-ch) //36
	fmt.Println(<-ch) //37
}
