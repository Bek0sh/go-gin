package service

import (
	"fmt"
	"project2/pkg/config"
	"project2/pkg/models"
	irepo "project2/pkg/repository/repoInterfaces"
	iservice "project2/pkg/service/serviceInterfaces"
	"project2/pkg/utils"
)

type AuthService struct {
	repo irepo.UserInterface
}

func NewAuthService(repo irepo.UserInterface) iservice.AuthServiceInterface {
	return &AuthService{repo: repo}
}

func (service *AuthService) Register(userInput models.RegisterUser) (*models.ResponseUser, error) {
	hashedPassword, err := utils.HashPassword(userInput.Password)

	if err != nil || userInput.ConfirmPassord != userInput.Password {
		return &models.ResponseUser{}, fmt.Errorf("something wrong with your password, error: %s", err.Error())
	}

	createUser := &models.User{
		Name:     userInput.Name,
		Surname:  userInput.Surname,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	createdUser, err := service.repo.CreateUser(createUser)

	if err != nil {
		return &models.ResponseUser{}, fmt.Errorf("failed to create user, error: %s", err.Error())
	}

	return createdUser, nil
}

func (service *AuthService) SignIn(userInput models.SignInUser) (string, string, error) {
	user, err := service.repo.FindUserByEmail(userInput.Email)

	if err != nil {
		return "", "", fmt.Errorf("Failed to find user with email=%s, error: %s", userInput.Email, err.Error())
	}

	if err = utils.CheckPassword(user.Password, userInput.Password); err != nil {
		fmt.Print("incorrect password")
		return "", "", fmt.Errorf("password is not correct, error: %s", err.Error())
	}

	config, _ := config.LoadConfig(".")

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, user.Email, config.AccessTokenPrivateKey)

	if err != nil {
		fmt.Print("failed to create access_token")

		return "", "", fmt.Errorf("failed to create acccess_token, error: %s", err.Error())
	}

	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, user.Email, config.RefreshTokenPrivateKey)

	if err != nil {
		fmt.Print("failed to create access_token")

		return "", "", fmt.Errorf("failed to create refresh_token, error: %s", err.Error())
	}

	return accessToken, refreshToken, nil
}
