package main

import "fmt"

func divide(a, b float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Cannot divide by zero!")
	}

	return a / b
}

func main() {
	result := divide(10, 2)
	fmt.Println("Result:", result)

	// This will cause a panic, but it will be recovered in the function.
	result = divide(5, 0)
	fmt.Println("Result:", result) // This line will still be executed.
}
