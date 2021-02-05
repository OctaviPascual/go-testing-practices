package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTemporaryFile(t *testing.T) *os.File {
	file, err := ioutil.TempFile("", "*.txt")
	require.NoError(t, err)

	t.Cleanup(func() {
		err := file.Close()
		require.NoError(t, err)
	})

	return file
}

func setTimeAt(t *testing.T, greeter *greeter, timeValue string) {
	greeter.now = func() time.Time {
		now, err := time.Parse(time.Kitchen, timeValue)
		require.NoError(t, err)
		return now
	}
}

func TestGreeterShould(t *testing.T) {
	t.Run("work with a file", func(t *testing.T) {
		file := newTemporaryFile(t)
		greeter := newGreeter(file)
		setTimeAt(t, greeter, "3:59PM")

		err := greeter.greet()
		require.NoError(t, err)

		b, err := ioutil.ReadFile(file.Name())
		require.NoError(t, err)
		assert.Equal(t, "Hello world at 3:59PM!", string(b))
	})

	t.Run("work with a buffer", func(t *testing.T) {
		var buffer bytes.Buffer
		greeter := newGreeter(&buffer)
		setTimeAt(t, greeter, "11:15AM")

		err := greeter.greet()
		require.NoError(t, err)

		assert.Equal(t, "Hello world at 11:15AM!", buffer.String())
	})

	t.Run("greet twice", func(t *testing.T) {
		file := newTemporaryFile(t)
		greeter := newGreeter(file)

		setTimeAt(t, greeter, "11:15AM")
		err := greeter.greet()
		require.NoError(t, err)

		setTimeAt(t, greeter, "11:16AM")
		err = greeter.greet()
		require.NoError(t, err)

		b, err := ioutil.ReadFile(file.Name())
		require.NoError(t, err)
		assert.Equal(t, "Hello world at 11:15AM!Hello world at 11:16AM!", string(b))
	})
}
