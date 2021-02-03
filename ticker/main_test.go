package main

import (
	"bytes"
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type safeBuffer struct {
	mutex  sync.Mutex
	buffer bytes.Buffer
}

func (b *safeBuffer) WriteString(s string) (int, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.WriteString(s)
}

func (b *safeBuffer) String() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.String()
}

func Test(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var buffer safeBuffer
	greeter := newGreeter(ctx, &buffer, time.Hour)

	tickerC := make(chan time.Time)
	greeter.tickerC = tickerC

	greeter.run()

	tickerC <- time.Now()
	f := func() bool { return assert.ObjectsAreEqual("Hello world!\n", buffer.String()) }
	assert.Eventually(t, f, time.Second, 10*time.Millisecond)

	tickerC <- time.Now()
	f = func() bool { return assert.ObjectsAreEqual("Hello world!\nHello world!\n", buffer.String()) }
	assert.Eventually(t, f, time.Second, 10*time.Millisecond)
}
