package main

import (
	"fmt"
)

type user struct {
	ID   int64
	name string
}

type userRepository interface {
	FindAll() ([]*user, error)
}

type staticUserRepository struct{}

func (staticUserRepository) FindAll() ([]*user, error) {
	return []*user{
		{ID: 1, name: "alice"},
		{ID: 2, name: "bob"},
		{ID: 3, name: "charlie"},
	}, nil
}

type userFinder struct {
	repository userRepository
}

func newUserFinder(repository userRepository) *userFinder {
	return &userFinder{
		repository: repository,
	}
}

func (uf *userFinder) find(ID int64) (*user, error) {
	allUsers, err := uf.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("could not find users: %w", err)
	}

	for _, user := range allUsers {
		if user.ID == ID {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with ID %d doesn't exist", ID)
}

func main() {
	userFinder := newUserFinder(staticUserRepository{})
	for i := int64(1); i < 5; i++ {
		user, err := userFinder.find(i)
		if err != nil {
			fmt.Printf("could not find user: %s\n", err)
			continue
		}
		fmt.Printf("found user with ID %d: %+v\n", i, user)
	}

}
