package main

import "fmt"
//给自定义类型加方法
type myInt int

func (i myInt) hello(){
	fmt.Println("这是一个int方法")
}

func main(){
	m:=myInt(100)
	m.hello()
}