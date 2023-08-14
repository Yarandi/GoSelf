package main

import (
	"fmt"
	"time"
)

func main() {

	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Counter", counter)
}
