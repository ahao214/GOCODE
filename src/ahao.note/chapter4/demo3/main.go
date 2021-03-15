package main

import (
	"fmt"
)

//匿名函数
func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}
