package main

import "fmt"

func main(){
	// 变量
	var i int = 10
	fmt.Println("i = ", i)

	// 类型推导
	var num = 10
	fmt.Println("num = ", num)

	name := "Tome"  //冒号不能省略
	fmt.Println("name is ",name)

	var one,two,three int
	fmt.Println("one=",one,"two=",two,"three=",three)

	var n1,myname,n3 = 100,"tom",888
	fmt.Println("n1=",n1,"name=",myname,"n3=",n3)
}