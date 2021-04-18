package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

//GORM更新操作
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

	var user User
	db.First(&user)

	user.Name = "aho"
	user.Age = 19
	//默认修改所有字段
	db.Save(&user)

	//更新指定的字段
	db.Model(&user).Update("name", "小王子")

	m1 := map[string]interface{}{
		"name":   "jack",
		"age":    29,
		"active": true,
	}
	//更新m1中所有字段
	db.Model(&user).Updates(m1)
	//只更新age字段
	db.Model(&user).Select("age").Updates(m1)
	//排除active字段，更新其他字段
	db.Model(&user).Omit("active").Updates(m1)

}
