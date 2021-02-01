package main

import (
	"fmt"
	"log"
)

type user struct {
	id   int64
	name string
}

func newUser(id int64, name string) (*user, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user id cannot be negative")
	}

	if name == "" {
		return nil, fmt.Errorf("user name cannot be empty")
	}

	return &user{
		id:   id,
		name: name,
	}, nil
}

type userRepository interface {
	save(user *user) error
}

type fakeUserRepository struct{}

func (*fakeUserRepository) save(user *user) error {
	fmt.Printf("saved user %+v", user)
	return nil
}

type userCreator struct {
	repository userRepository
}

func newUserCreator(repository userRepository) *userCreator {
	return &userCreator{
		repository: repository,
	}
}

func (uc *userCreator) create(id int64, name string) error {
	user, err := newUser(id, name)
	if err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	return uc.repository.save(user)
}

func main() {
	userCreator := newUserCreator(&fakeUserRepository{})
	err := userCreator.create(1, "John")
	if err != nil {
		log.Fatalf("could not create user: %s", err)
	}
}
