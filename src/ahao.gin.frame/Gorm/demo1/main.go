package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"` //跟Name字段设置默认值为小王子
	Age  int64
}

//GORM创建记录及字段
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
	u := User{
		Name: "aho",
		Age:  18,
	}
	//判断主键是否为空
	db.NewRecord(&u)

	db.Create(&u)

}
