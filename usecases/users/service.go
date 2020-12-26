package users

import (
	"home.com/lists/backend/entities"
	"home.com/lists/backend/repositories"
)

type Service struct {
	Repository repositories.UsersRepository
}

func NewService(r repositories.UsersRepository) *Service {
	return &Service{
		Repository: r,
	}
}

func (service *Service) ListUsers() ([]*entities.User, error) {
	return service.Repository.ListUsers()
}

func (service *Service) CreateUser(email, password, first_name, last_name string) (*entities.User, error) {

	isEmailAlreadyUsed, err := service.Repository.DoesUserEmailExists(email)

	if err != nil {
		return nil, err
	}

	if isEmailAlreadyUsed {
		return nil, entities.ErrEmailAlreadyExists
	}

	newUser, err := entities.NewUser(email, password, first_name, last_name)

	if err != nil {
		return nil, err
	}

	result, err := service.Repository.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	newUser.ID = result

	return newUser, nil
}
