package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	age  int
}

func demo(ce []Student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 99
	ce[1].name = "小王"
}

func main() {
	var c []Student //定义一个切片类型的结构体
	c = []Student{
		Student{1, "小明", 22},
		Student{2, "小米", 33},
	}
	fmt.Println("执行demo之前的值：", c)
	demo(c)
	fmt.Println("执行demo之后的值：", c)

	//删除map类型的结构体
	s := make(map[int]Student)
	s[1] = Student{1, "小明", 22}
	s[2] = Student{2, "小王", 33}
	fmt.Println("删除之前：", s)
	delete(s, 2)
	fmt.Println("删除之后：", s)

}
