package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS \t\t", runtime.GOOS)
	fmt.Println("ARCH \t\t", runtime.GOARCH)
	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()

	bar()

	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println("World")
	}
}
