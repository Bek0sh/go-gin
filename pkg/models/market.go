package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name    string `json:"name" gorm:"type: varchar(60); not null"`
	Price   int    `json:"price" gorm:"type: int; not null"`
	User_ID uint   `json:"user_id" sql:"REFERENCES users(ID)"`
	User    User   `gorm:"foreignKey:User_ID"`
}

type ProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
