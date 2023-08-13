package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

// my brain structure
func main() {
	DoWork()
	w.Wait()
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
	go CheckBalance()
}

func CheckBalance() {
	fmt.Println("How?")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("go and f... yourself!")
	w.Done()
}
