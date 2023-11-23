package domain

import (
	"errors"
	"log"
)

var (
	ErrEmptyName = errors.New("Empty name supplied")
	ErrEmptyEmail = errors.New("Empty email supplied")
)

type FirstName string
type LastName string
type Email string
type IsAuthor bool

type User struct {
	FirstName FirstName
	LastName  LastName
	Email     Email
	IsAuthor  IsAuthor
}

func NewUser(firstName string, lastName string, email string, isAuthor bool) (User, error) {
	if firstName == "" {
		log.Print(ErrEmptyName)
		return User{}, ErrEmptyName
	}
	
	if lastName == "" {
		log.Print(ErrEmptyName)
		return User{}, ErrEmptyName
	}

	if email == "" {
		log.Print(ErrEmptyEmail)
		return User{}, ErrEmptyEmail
	}

	return User{
		FirstName(firstName),
		LastName(lastName),
		Email(email),
		IsAuthor(isAuthor),
	}, nil
}