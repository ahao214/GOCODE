package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  string
}

type Student struct {
	Person
	id   int
	addr string
	//同名字段
	name string
}

//同名字段
func main() {
	var s Student
	// 给自己字段赋值
	s.name = "mh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下:
	s.Person.name = "枯藤"
	fmt.Println(s)
}
