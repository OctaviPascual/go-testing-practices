package main

import (
	"io"
	"log"
	"os"
)

// Greeter writes greeting messages
type Greeter struct {
	writer io.StringWriter
}

// NewGreeter returns a new greeter
func NewGreeter(writer io.StringWriter) *Greeter {
	return &Greeter{
		writer: writer,
	}
}

// Greet writes a greeting message to the underlying writer
func (g *Greeter) Greet() error {
	_, err := g.writer.WriteString("Hello world!")
	return err
}

func main() {
	greeter := NewGreeter(os.Stdout)
	err := greeter.Greet()
	if err != nil {
		log.Fatalf("could not greet: %s", err)
	}
}
