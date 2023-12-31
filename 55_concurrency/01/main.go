package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var w sync.WaitGroup
var w1 sync.WaitGroup

// my brain structure
func main() {

	fmt.Println("OS \t\t", runtime.GOOS)
	fmt.Println("ARCH \t\t", runtime.GOARCH)
	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	DoWork()

	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	w.Wait()
	w1.Wait()

}

func DoWork() {
	fmt.Println("Be a good boy, don't lose your focusing!!!")
	w.Add(1)
	go Offroad()
}

func Offroad() {
	fmt.Println("What if I can buy a Arb Old Man Emu 3 Inch Lift Kit for my car???")
	for i := 0; i < 5; i++ {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("Think... \t")
	}
	fmt.Println("How can i effort that??")
	w1.Add(1)
	go CheckBalance()
}

func CheckBalance() {
	fmt.Println("How?")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("go and f... yourself!")
	w.Done()
	w1.Done()
}
