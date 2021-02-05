package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubUserRepository struct {
	// use this function to override stub default implementation
	findAll func() ([]*user, error)
}

func (f stubUserRepository) FindAll() ([]*user, error) {
	if f.findAll != nil {
		return f.findAll()
	}

	return []*user{
		{ID: 1, name: "alice"},
		{ID: 2, name: "bob"},
	}, nil
}

func TestUserFinderFindShould(t *testing.T) {
	t.Run("successfully find an existing user", func(t *testing.T) {
		userFinder := newUserFinder(stubUserRepository{})

		user, err := userFinder.find(1)
		require.NoError(t, err)

		assert.Equal(t, "alice", user.name)
	})

	t.Run("return an error when repository returns an error", func(t *testing.T) {
		userFinder := newUserFinder(stubUserRepository{
			findAll: func() ([]*user, error) {
				return nil, errors.New("always fail")
			},
		})

		_, err := userFinder.find(1)
		require.Error(t, err)
	})

	t.Run("return an error when user doesn't exist", func(t *testing.T) {
		userFinder := newUserFinder(stubUserRepository{})

		_, err := userFinder.find(3)
		require.Error(t, err)
	})
}
