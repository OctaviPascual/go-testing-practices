package main

import "fmt"

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = fizz + buzz
)

// FizzBuzz formats an integer according to the following rules:
// - if the number is divisible by 3 then return Fizz
// - if the number is divisible by 5 then return Buzz
// - if the number is divisible by 3 and 5 then return FizzBuzz
// - otherwise return the integer as it is
func FizzBuzz(i int) string {
	switch {
	case i%15 == 0:
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
