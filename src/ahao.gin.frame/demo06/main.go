package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzun/gorm"
	"shangguigu.com/Method/demo11/model"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100)"`
}

//修改表的名称
func (c User) TableName() string {
	return "aho_users"
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	ClassID     uint
	StudentName string
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers"`
	TeacherID   uint
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	StudentID   uint
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers"`
}

func main() {
	db, _ := gorm.Open("mysql", "root:Aa@123/ginclass?charset=utf8mb4&parseTime=True&loc=local")
	//创建表
	db.AutoMigrate(&Class{}, &Student{}, &Teacher{}, &IDCard{})
	i := IDCard{
		Num: 214,
	}
	s := Student{
		StudentName: "aho",
		IDCard:      i,
	}

	t := Teacher{
		TeacherName: "老师傅",
		Students:    []Student{s},
	}
	c := Class{
		ClassName: "214班",
		Students:  []Student{s},
	}

	_ = db.Create(&c).Error
	_ = db.Create(&t).Error

	defer db.Close()
}
