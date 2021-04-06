package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzun/gorm"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	db, err := db.Open("mysql", "root:Aa@123/ginclass?charset=utf8mb4&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	//创建表
	db.AutoMigrate(&HelloWorld{})

	//新增数据
	db.Create(&HelloWorld{
		Name: "aho",
		Age:  19,
		Sex:  true,
	})
	db.Create(&HelloWorld{
		Name: "枯藤",
		Age:  1000,
		Sex:  false,
	})

	//查询数据
	var hello HelloWorld
	db.First(&hello, "name=?", "aho")
	fmt.Println(hello)
	db.Find(&hello, "age < ?", 29)
	fmt.Println(hello)

	db.Where("name=?", "aho").Find(&hello)

	//修改数据
	db.Where("id=?", 1).First(&HelloWorld{}).Updates("name", "不二")
	db.Where("id=?", 1).First(&HelloWorld{}).Updates(HelloWorld{
		Name: "不二",
		Age:  90,
	})

	//删除数据
	//逻辑删除
	db.Where("id in (?)", []int{1, 2}).Delete(&HelloWorld{})

	//物理删除
	db.Where("id in (?)", []int{1, 2}).Unscoped().Delete(&HelloWorld{})

	defer db.Close()

}
