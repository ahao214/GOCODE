package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好GO语言！\n", p.name)
}

func main() {
	p1 := NewPerson("测试", 25)
	p1.Dream()

}
