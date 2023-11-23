package ports

import (
	"bythecover/backend/internal/core/domain"
	"context"
	"database/sql"
	"time"
)

type UserRepo interface {
	Save(domain.User) error
	GetAll() ([]UserResp, error)
	GetUser(int, context.Context) (UserResp, error)
}

var (
	ErrUserNotFound = sql.ErrNoRows
)

type UserService interface {
	Create(UserResp) error
	GetAll() ([]UserResp, error)
	GetUser(int, context.Context) (UserResp, error)
}

type UserResp struct {
	Id        *int     	`json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsAuthor  bool      `json:"is_author"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type UserReq struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsAuthor  bool      `json:"is_author"`
}