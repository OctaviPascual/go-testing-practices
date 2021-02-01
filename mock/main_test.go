package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) save(user *user) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) expectUserAndSucceed(user *user) {
	m.On("save", user).Return(nil).Once()
}

func (m *mockUserRepository) expectUserAndFail(user *user) {
	m.On("save", user).Return(errors.New("error saving user")).Once()
}

func TestUserCreatorCreateShould(t *testing.T) {
	t.Run("successfully create a user", func(t *testing.T) {
		mockUserRepository := &mockUserRepository{}
		mockUserRepository.expectUserAndSucceed(&user{id: 2, name: "John"})
		userCreator := newUserCreator(mockUserRepository)

		err := userCreator.create(2, "John")
		require.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("return an error when repository fails to create a user", func(t *testing.T) {
		mockUserRepository := &mockUserRepository{}
		mockUserRepository.expectUserAndFail(&user{id: 4, name: "James"})
		userCreator := newUserCreator(mockUserRepository)

		err := userCreator.create(4, "James")
		require.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("return an error when user is invalid", func(t *testing.T) {
		mockUserRepository := &mockUserRepository{}
		userCreator := newUserCreator(mockUserRepository)

		err := userCreator.create(2, "")
		require.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}
