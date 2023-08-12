package interfaces

import "project2/pkg/models"

type AuthServiceInterface interface {
	Register(userInput models.RegisterUser) (*models.ResponseUser, error)
	SignIn(userInput models.SignInUser) (string, string, error)
}
