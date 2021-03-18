package main

import (
	"fmt"
)

func main() {
	//通过数组创建切片
	var arrary = [...]int{0, 1, 2, 3, 4, 5, 6}
	one := arrary[0:4]
	two := arrary[:4]
	three := arrary[2:]
	fmt.Printf("%v\n", one)
	fmt.Printf("%v\n", two)
	fmt.Printf("%v\n", three)

	//通过内置函数make创建切片
	a := make([]int, 10)
	b := make([]int, 10, 15)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}
