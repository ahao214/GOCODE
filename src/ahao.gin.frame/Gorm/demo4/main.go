package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

//GORM删除操作
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
	db.Create(&u1)

	//删除-删除是软删除
	var u = UserInfo{}
	u.Name = "aho"
	db.Delete(&u)

	db.Where("name = ?", "aho").Delete(UserInfo{})
	db.Where(UserInfo{}, "Age = ?", 19)

	//物理删除
	db.Unscoped().Where("name = ?", "aho").Delete(UserInfo{})
}
