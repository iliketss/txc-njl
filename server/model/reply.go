package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	CommentId int
	Content   string
	CUserId   int
	RUserId   int
	User      []User `gorm:"many2many:User_Reply;"`
}

/// a: wtf     c_id  g_id u_id content
/// b->a addd  c_id  g_id r_id u_id content
/// a->b wdnmd
