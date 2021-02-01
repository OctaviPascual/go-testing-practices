package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

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

func TestGreeterShould(t *testing.T) {
	t.Run("work with a file", func(t *testing.T) {
		file := newTemporaryFile(t)
		greeter := NewGreeter(file)

		err := greeter.Greet()
		require.NoError(t, err)

		b, err := ioutil.ReadFile(file.Name())
		require.NoError(t, err)
		assert.Equal(t, "Hello world!", string(b))
	})

	t.Run("work with a buffer", func(t *testing.T) {
		var buffer bytes.Buffer
		greeter := NewGreeter(&buffer)

		err := greeter.Greet()
		require.NoError(t, err)

		assert.Equal(t, "Hello world!", buffer.String())
	})

	t.Run("greet twice", func(t *testing.T) {
		file := newTemporaryFile(t)
		greeter := NewGreeter(file)

		err := greeter.Greet()
		require.NoError(t, err)
		err = greeter.Greet()
		require.NoError(t, err)

		b, err := ioutil.ReadFile(file.Name())
		require.NoError(t, err)
		assert.Equal(t, "Hello world!Hello world!", string(b))
	})
}
