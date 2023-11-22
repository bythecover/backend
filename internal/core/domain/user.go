package domain

import "time"

type User struct {
	Id        int32     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsAuthor  bool      `json:"is_author"`
	CreatedAt time.Time `json:"created_at"`
}
