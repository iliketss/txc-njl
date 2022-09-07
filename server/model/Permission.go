package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ID     int
	PerUrl string
	User   User `gorm:"embedded"`
}
