package main

import(
	"fmt"
)

//延迟语句

var i=0
func print(i int){
	fmt.Println(i)
}

func main(){
	defer fmt.Println("world")
	fmt.Println("hello")

	//defer其实是一个栈，遵循先入后出
	for ;i<5;i++{
		defer print(i)
	}

}