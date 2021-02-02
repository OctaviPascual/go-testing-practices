package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type user struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func Test(t *testing.T) {
	file, err := os.Open(filepath.Join("testdata", "user.json"))
	require.NoError(t, err)
	defer func() {
		err := file.Close()
		require.NoError(t, err)
	}()

	b, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	var user user
	err = json.Unmarshal(b, &user)
	require.NoError(t, err)

	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "James", user.Name)
}
