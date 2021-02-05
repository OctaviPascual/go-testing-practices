package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type spyUserRepository struct {
	saveCalled bool
}

func (m *spyUserRepository) save(_ *user) error {
	m.saveCalled = true
	return nil
}

func TestUserCreatorCreateShould(t *testing.T) {
	t.Run("successfully create a user", func(t *testing.T) {
		spyUserRepository := &spyUserRepository{}
		userCreator := newUserCreator(spyUserRepository)

		err := userCreator.create(2, "John")
		require.NoError(t, err)

		assert.True(t, spyUserRepository.saveCalled)
	})

	t.Run("return an error when user is invalid", func(t *testing.T) {
		spyUserRepository := &spyUserRepository{}
		userCreator := newUserCreator(spyUserRepository)

		err := userCreator.create(2, "")
		require.Error(t, err)

		assert.False(t, spyUserRepository.saveCalled)
	})
}
