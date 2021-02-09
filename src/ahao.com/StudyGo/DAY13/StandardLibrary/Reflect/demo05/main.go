package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//匿名字段
type Boy struct {
	User
	Addr string
}

//匿名字段
func main() {
	m := Boy{User{1, "kuteng", 20}, "sh"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Printf("%#v\n", t.Field(0))
	//值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))

}
