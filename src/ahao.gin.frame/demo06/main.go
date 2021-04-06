package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzun/gorm"
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
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(student)
	})

	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").First(&student, "id=?", id)
		c.JSON(200, gin.H{
			"s": student,
		})
	})

	r.GET("/class/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		_ = c.BindJSON(&class)
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").First(&class, "id=?", id)
		c.JSON(200, gin.H{
			"c": class,
		})

	})

	r.Run(":8080")
}
