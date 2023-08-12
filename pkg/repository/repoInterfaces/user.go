package interfaces

import "project2/pkg/models"

type UserInterface interface {
	CreateUser(*models.User) (*models.ResponseUser, error)
	SignInUser(*models.SignInUser) error
	FindUserById(id int) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}
