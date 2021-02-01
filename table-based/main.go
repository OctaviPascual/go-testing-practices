package main

import "fmt"

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = fizz + buzz
)

func FizzBuzz(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return fizzBuzz
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	default:
		return fmt.Sprintf("%d", i)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
