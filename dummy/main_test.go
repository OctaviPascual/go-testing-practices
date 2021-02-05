package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type dummyUserRepository struct{}

func (m *dummyUserRepository) save(_ *user) error {
	panic("should not be called")
}

func TestUserCreatorCreateShould(t *testing.T) {
	t.Run("return an error when user ID is negative", func(t *testing.T) {
		userCreator := newUserCreator(&dummyUserRepository{})

		err := userCreator.create(-2, "")
		require.Error(t, err)
	})

	t.Run("return an error when user name is empty", func(t *testing.T) {
		userCreator := newUserCreator(&dummyUserRepository{})

		err := userCreator.create(2, "")
		require.Error(t, err)
	})
}
