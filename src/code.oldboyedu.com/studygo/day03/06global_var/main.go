package main

import "fmt"

//定义一个全局变量
var x=100

func f1(){
	fmt.Println(x)
}
func main(){
	f1()
}