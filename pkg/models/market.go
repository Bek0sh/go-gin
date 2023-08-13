package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"name" gorm:"type: varchar(60); not null"`
	Price   float32 `json:"price" gorm:"type:numeric(8,2); not null"`
	User_Id uint    `json:"user_id" gorm:"type:"`
}
