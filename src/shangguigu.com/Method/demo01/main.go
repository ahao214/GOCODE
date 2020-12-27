package main

import(
	"fmt"
)

type Person struct{
	Name string
}

//给Person类型绑定一个方法
func(p Person) test(){
	p.Name="Jack"
	fmt.Println("test()的名字是：",p.Name)
}

//方法
func main(){
	var p Person
	p.Name="Tom"
	p.test()
	fmt.Println("main() p.Name是：",p.Name)
}