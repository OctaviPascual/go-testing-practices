package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibShould(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"return 0 when input is 0": {
			input:    0,
			expected: 0,
		},
		"return 55 when input is 10": {
			input:    10,
			expected: 55,
		},
		"return 5 when input is 5": {
			input:    5,
			expected: 5,
		},
		"return 6765 when input is 20": {
			input:    20,
			expected: 6765,
		},
		"return 433494437 when input is 43": {
			input:    43,
			expected: 433494437,
		},
		"return 1134903170 when input is 45": {
			input:    45,
			expected: 1134903170,
		},
	}

	for name, test := range tests {
		// hide the loop variable with a local one
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fib(test.input))
		})
	}
}
