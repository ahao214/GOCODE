package main

import (
	"fmt"
)

func xor(x, y int) int {
	return (x | y) & (^x | ^y)
}

//如何不使用^操作实现异或运算
func main() {
	fmt.Println("方法一：")
	fmt.Println(xor(3, 5))
}
