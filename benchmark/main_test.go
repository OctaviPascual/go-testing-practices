package main

import (
	"testing"
)

var result int

func BenchmarkFib(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		// always record the result of Fib to prevent the compiler eliminating the function call
		r = fib(10)
	}
	// always store the result to a package level variable so the compiler cannot eliminate the Benchmark itself
	result = r
}
