package main

import "fmt"

func someFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	// Code that might panic
	panic("Something went wrong!")
}

func main() {
	someFunction()
	fmt.Println("After panic")
}
