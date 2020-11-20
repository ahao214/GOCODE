package main

import "fmt"

//用结构体模拟继承

type animal struct{
	name string
}

//给animal实现一个移动的方法
func(a animal) move(){
	fmt.Printf("%s 谁会动",a.name)
}

type dog struct{
	feet uint8
	animal
}

//给dog实现一个汪汪汪的方法
func(d dog)wang(){
	fmt.Printf("%s在汪汪汪~",d.name)
}

func main(){
	d:=dog{
		animal:animal{
			name:"Tom",
		},
		feet:4,
	}
	fmt.Println(d)
	d.wang()
	d.move()
}