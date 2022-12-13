package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllColors(t *testing.T) {
	// Ensure that when a new color is added, at least this unit test will fail.
	// Make sure to update all unit tests of this file when adding a new color.
	require.Len(t, AllColors, 3)

	assert.Contains(t, AllColors, Red)
	assert.Contains(t, AllColors, Green)
	assert.Contains(t, AllColors, Blue)
}

func TestString(t *testing.T) {
	assert.Equal(t, "red", Red.String())
	assert.Equal(t, "green", Green.String())
	assert.Equal(t, "blue", Blue.String())
}
