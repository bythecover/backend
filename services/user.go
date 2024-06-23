package services

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"
)

type userService struct {
	userRepo persistence.UserRepo
}

type UserService interface {
	Create(model.UserResp) error
	GetAll() ([]model.UserResp, error)
	GetUser(int) (model.UserResp, error)
}

func NewUserService(userRepo persistence.UserRepo) userService {
	return userService{
		userRepo,
	}
}

func (service userService) Create(input model.UserResp) error {
	user, err := model.NewUser(input.Id, input.FirstName, input.LastName, input.Email, input.IsAuthor)

	if err != nil {
		return err
	}

	err = service.userRepo.Save(user)

	return err
}

func (service userService) GetAll() ([]model.UserResp, error) {
	users, err := service.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service userService) GetUser(id int) (model.UserResp, error) {
	return service.userRepo.GetUser(id)
}
