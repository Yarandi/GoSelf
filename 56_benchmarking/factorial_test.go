package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10) // Benchmark the Factorial function with 10
	}
}
