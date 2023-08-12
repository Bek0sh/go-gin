package repository

import (
	"fmt"
	"project2/pkg/models"
	interfaces "project2/pkg/repository/repoInterfaces"
	"project2/pkg/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.UserInterface {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *models.User) (*models.ResponseUser, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return &models.ResponseUser{}, err
	}

	output := &models.ResponseUser{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}

	return output, nil
}

func (repo *userRepository) SignInUser(user *models.SignInUser) error {
	var password string

	if err := repo.db.First(&password, "email=?", user.Email); err != nil {
		return fmt.Errorf("Failed to find user with email %s", user.Email)
	}

	err := utils.CheckPassword(password, user.Password)

	if err != nil {
		return fmt.Errorf("password does not match")
	}

	return nil
}

func (repo *userRepository) FindUserById(id int) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, "id=?", id).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (repo *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, "email=?", email).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
