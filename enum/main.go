package main

import (
	"fmt"
)

// Color represents a color
// We encapsulate enums in struct for extra safety (cf https://threedots.tech/post/safer-enums-in-go/)
type Color struct {
	name string
}

var (
	// Red represents the red color
	Red = Color{name: "red"}
	// Green represents the blue color
	Green = Color{name: "green"}
	// Blue represents the blue color
	Blue = Color{name: "blue"}
)

// AllColors contains all the colors
var AllColors = []Color{
	Red,
	Green,
	Blue,
}

// String returns the color as a string
func (c Color) String() string {
	return c.name
}

func main() {
	fmt.Printf("We have 3 enums: %s, %s and %s\n", Red, Green, Blue)
}
