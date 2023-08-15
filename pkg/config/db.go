package config

import (
	"log"
	"project2/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn := "postgres://postgres:1234@localhost:5432/gogin?sslmode=disable"

	data, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to db")
	}

	DB = data

	DB.AutoMigrate(models.User{}, models.Product{})
}
