package main

import (
	"context"
	"io"
	"os"
	"time"
)

// greeter writes greeting messages at every given interval
type greeter struct {
	writer io.StringWriter

	ticker  *time.Ticker
	tickerC <-chan time.Time
}

func newGreeter(writer io.StringWriter, interval time.Duration) *greeter {
	ticker := time.NewTicker(interval)

	return &greeter{
		writer:  writer,
		ticker:  ticker,
		tickerC: ticker.C,
	}
}

func (g *greeter) run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-g.tickerC:
				_, _ = g.writer.WriteString("Hello world!\n")
			case <-ctx.Done():
				g.ticker.Stop()
				return
			}
		}
	}()
}

func main() {
	greeter := newGreeter(os.Stdout, time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	greeter.run(ctx)

	<-ctx.Done()
}
