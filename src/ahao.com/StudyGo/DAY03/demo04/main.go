package main

import (
	"fmt"
)

//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//Class 班级结构体
type Class struct {
	ClassName string
	Grade     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
	Class   //匿名结构体
}

func main() {

	user1 := User{
		Name:   "pprof",
		Gender: "男",
		Address: Address{
			Province: "上海",
			City:     "上海",
		},
		Class: Class{
			ClassName: "一班",
			Grade:     "一年级",
		},
	}

	fmt.Printf("user1=%#v\n", user1)

}
