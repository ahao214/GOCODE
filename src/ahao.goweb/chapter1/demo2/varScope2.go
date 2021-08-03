package main

import (
	"fmt"
)

//声明全局变量
var global int = 8

func main() {
	//声明局部变量
	var global int = 999
	fmt.Printf("global=%d\n", global)
}
