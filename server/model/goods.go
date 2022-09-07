package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name    string    `form:"name";gorm:"unique;not null"`
	Price   float32   `form:"price"`
	Comment []Comment `gorm:"many2many:Goods_Comment;"`
}
