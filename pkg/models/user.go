package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Surname  string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255); not null; uniqueIndex"`
	Password string `gorm:"not null"`
}

type RegisterUser struct {
	Name           string `json:"name" binding:"required"`
	Surname        string `json:"surname" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ConfirmPassord string `json:"confirm_password" binding:"required"`
}

type ResponseUser struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type SignInUser struct {
	Email    string
	Password string
}
