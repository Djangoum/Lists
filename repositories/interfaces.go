package repositories

import "home.com/lists/backend/entities"

type UsersRepository interface {
	ListUsers() ([]*entities.User, error)
	DoesUserEmailExists(email string) (bool, error)
	CreateUser(user *entities.User) (int64, error)
}
