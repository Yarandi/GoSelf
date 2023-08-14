package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1) // Limiting to one processor for demonstration
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go printNumbers()
	go printLetters()

	wg.Wait()

	fmt.Println("\nEnd Goroutines")
}

func printNumbers() {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		runtime.Gosched() // Yield processor to other goroutines
	}
}

func printLetters() {
	defer wg.Done()

	for char := 'A'; char < 'A'+5; char++ {
		fmt.Printf("%c ", char)
		runtime.Gosched() // Yield processor to other goroutines
	}
}
