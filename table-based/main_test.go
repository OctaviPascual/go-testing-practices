package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzShould(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected string
	}{
		"return Fizz when input is 3": {
			input:    3,
			expected: "Fizz",
		},
		"return Buzz when input is 100": {
			input:    100,
			expected: "Buzz",
		},
		"return 1 when input is 1": {
			input:    1,
			expected: "1",
		},
		"return FizzBuzz when input is 30": {
			input:    30,
			expected: "FizzBuzz",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, FizzBuzz(test.input))
		})
	}
}
