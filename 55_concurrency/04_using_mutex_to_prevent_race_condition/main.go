package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {

	var counter int
	for i := 0; i < 1000; i++ {

		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()

	}
	time.Sleep(time.Second)
	fmt.Println("Counter", counter)
}
