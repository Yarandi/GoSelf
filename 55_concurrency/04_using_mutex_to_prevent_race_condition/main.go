package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	defer wg.Wait()
	var counter int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter", counter)
}
