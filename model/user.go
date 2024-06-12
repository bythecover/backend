package model

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

var (
	ErrEmptyName    = errors.New("Empty name supplied")
	ErrEmptyEmail   = errors.New("Empty email supplied")
	ErrUserNotFound = sql.ErrNoRows
)

type UserResp struct {
	Id        *int       `json:"id,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	IsAuthor  bool       `json:"is_author"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type UserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsAuthor  bool   `json:"is_author"`
}

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
