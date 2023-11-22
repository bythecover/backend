package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
)

type userService struct {
	repo ports.UserRepo
}

func NewUserService(repo ports.UserRepo) userService {
	return userService{
		repo,
	}
}

func (service userService) Create(user domain.User) {
	service.repo.Save(user);
}

func (service userService) GetAll() ([]domain.User, error) {
	users, err := service.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
