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

func (service userService) Create(input ports.UserResp) error {
	user, err := domain.NewUser(input.FirstName, input.LastName, input.Email, input.IsAuthor)

	if err != nil {
		return err
	}

	err = service.repo.Save(user)

	return err
}

func (service userService) GetAll() ([]ports.UserResp, error) {
	users, err := service.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service userService) GetUser(id int) (ports.UserResp, error) {
	return service.repo.GetUser(id)
}
