package main

import (
	"context"
	"io"
	"os"
	"time"
)

// greeter writes greeting messages at every given interval
type greeter struct {
	ctx context.Context

	writer io.StringWriter

	ticker  *time.Ticker
	tickerC <-chan time.Time
}

func newGreeter(ctx context.Context, writer io.StringWriter, interval time.Duration) *greeter {
	ticker := time.NewTicker(interval)

	return &greeter{
		ctx:     ctx,
		writer:  writer,
		ticker:  ticker,
		tickerC: ticker.C,
	}
}

func (g *greeter) run() {
	go func() {
		for {
			select {
			case <-g.tickerC:
				_, _ = g.writer.WriteString("Hello world!\n")
			case <-g.ctx.Done():
				g.ticker.Stop()
				return
			}
		}
	}()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	greeter := newGreeter(ctx, os.Stdout, time.Second)
	greeter.run()

	<-ctx.Done()
}
