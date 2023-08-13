package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string       `json:"name" gorm:"type: varchar(60); not null"`
	Price float32      `json:"price" gorm:"type:numeric(8,2); not null"`
	User  ResponseUser `json:"user" gorm:"foreignKey:ID"`
}

type ProductInput struct {
	Name  string       `json:"name"`
	Price float32      `json:"price"`
	User  ResponseUser `json:"user"`
}

type ProductResponse struct {
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	UserName  string  `json:"user_name"`
	UserEmail string  `json:"user_email"`
}
