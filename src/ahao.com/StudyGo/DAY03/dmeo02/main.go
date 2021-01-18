package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "prof.cn", age: 19},
		{name: "测试", age: 23},
		{name: "博客", age: 24},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
