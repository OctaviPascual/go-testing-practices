package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// greeter writes greeting messages with current time
type greeter struct {
	writer io.StringWriter
	now    func() time.Time
}

func newGreeter(writer io.StringWriter) *greeter {
	return &greeter{
		writer: writer,
		now:    time.Now,
	}
}

func (g *greeter) greet() error {
	formattedNow := g.now().Format(time.Kitchen)
	_, err := g.writer.WriteString(fmt.Sprintf("Hello world at %s!", formattedNow))
	return err
}

func main() {
	greeter := newGreeter(os.Stdout)
	err := greeter.greet()
	if err != nil {
		log.Fatalf("could not greet: %s", err)
	}
}
