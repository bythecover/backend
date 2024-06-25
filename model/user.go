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
	Id        string     `json:"id,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
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
type Role string

type User struct {
	Id        string
	FirstName FirstName
	LastName  LastName
	Email     Email
	Role      Role
}

func NewUser(Id string, firstName string, lastName string, email string, role string) (User, error) {
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
		Id,
		FirstName(firstName),
		LastName(lastName),
		Email(email),
		Role(role),
	}, nil
}
