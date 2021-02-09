package main

import (
	"fmt"
	"reflect"
)

//定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello:", name)
}

//调用方法
func main() {
	u := User{1, "kuteng", 20}
	v := reflect.ValueOf(u)
	//获取方法
	m := v.MethodByName("Hello")
	//构建一些参数
	args := []reflect.Value{reflect.ValueOf("6666")}
	m.Call(args)
}
