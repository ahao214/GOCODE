package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

//GORM查询操作
func main() {
	//连接数据库
	db, err := grom.Open("mysql", "root:root@123(127.0.0.1:12306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//创建
	u1 := User{
		Name: "aho",
		Age:  30,
	}
	db.Create(u1)

	u2 := User{
		Name: "chen",
		Age:  23,
	}
	db.Create(u2)

	//查询
	var user User
	//根据主键查询第一条记录
	db.First(&user)
	//随机获取一条记录
	db.Take(&user)
	//根据主键查询最后一条记录
	db.Last(&user)

	//查询所有记录
	var users []User
	db.Find(&users)

	db.Where("name = ?", "aho").First(&user)

	db.Where(&User{Name: "aho", Age: 19}).First(&user)

}
