package services

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"
)

type userService struct {
	repo persistence.UserRepo
}

type UserService interface {
	Create(model.UserResp) error
	GetAll() ([]model.UserResp, error)
	GetUser(int) (model.UserResp, error)
}

func NewUserService(repo persistence.UserRepo) userService {
	return userService{
		repo,
	}
}

func (service userService) Create(input model.UserResp) error {
	user, err := model.NewUser(input.FirstName, input.LastName, input.Email, input.IsAuthor)

	if err != nil {
		return err
	}

	err = service.repo.Save(user)

	return err
}

func (service userService) GetAll() ([]model.UserResp, error) {
	users, err := service.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service userService) GetUser(id int) (model.UserResp, error) {
	return service.repo.GetUser(id)
}
