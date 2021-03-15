package main

import "fmt"

//延迟调用
func main() {
	x, y := 1, 2
	defer func(a int) {
		fmt.Println("defer x,y=", a, y) //y为闭包引用
	}(x)
	x += 10
	y += 200
	fmt.Println(x, y)
}
