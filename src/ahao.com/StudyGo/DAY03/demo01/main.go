package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p person
	p.name = "Jack"
	p.city = "上海"
	p.age = 18
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%#v\n", p)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "prof.com"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//创建指针类型结构体
	//可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 19
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2)

	//去结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3)

	//使用键值对初始化结构体
	p5 := person{
		name: "pro.com",
		city: "北京",
		age:  10,
	}
	fmt.Printf("p5=%#v\n", p5)

	//使用值的列表初始化
	p6 := person{
		"pro.cn",
		"beijing",
		19,
	}
	fmt.Printf("p6=%#v\n", p6)

}
