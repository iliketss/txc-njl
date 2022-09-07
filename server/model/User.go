package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `form:"name";gorm:"unique;not null"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
}
