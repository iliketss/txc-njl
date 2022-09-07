package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	GoodsId int
	UserID  int
	User    []User  `gorm:"many2many:User_Comment;"`
	Reply   []Reply `gorm:"many2many:Comment_Reply;"`
}
