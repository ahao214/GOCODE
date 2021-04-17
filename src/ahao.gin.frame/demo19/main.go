package main

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

//gorm demo
func main() {
	//连接数据库
	db, err := grom.Open("mysql", "root:root@123(127.0.0.1:12306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表
	db.AutoMigrate(&UserInfo{})
	//创建数据行
	u1 := UserInfo{
		ID:     1,
		Name:   "aho",
		Gender: "男",
		Hobby:  "跑步",
	}
	db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u)

	//更新
	db.Model(&u).Update("Hobby", "双色球")

	//删除
	db.Delete(&u)
}
