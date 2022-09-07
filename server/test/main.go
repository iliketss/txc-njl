package main

import (
	"gorm.io/gorm"
	"txc-njl/server/DB"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student `gorm:"many2many:class_student;"`
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teacher;"`
}
type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:teacher_student;"`
}
type IDCard struct {
	gorm.Model
	StudentID uint
	Num       string
}

func main() {
	db := DB.InitDB()

	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
	i := IDCard{Num: "2018220243"}

	s := Student{StudentName: "唐秀财", IDCard: i}

	t := Teacher{TeacherName: "思思", Students: []Student{s}}
	c := Class{ClassName: "软件二班", Students: []Student{s}}

	db.Create(&c)
	db.Create(&t)
}
