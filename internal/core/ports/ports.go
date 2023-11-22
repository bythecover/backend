package ports

import "bythecover/backend/internal/core/domain"

type UserRepo interface {
	Save(domain.User) error
	GetAll() ([]domain.User, error)
}

type UserService interface {
	Create(domain.User)
	GetAll() ([]domain.User, error)
}
