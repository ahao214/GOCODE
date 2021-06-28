package main

import (
	"fmt"
)

//用++实现加法操作(限制条件：至少有一个非负数)
func add(a, b int) int {
	if a < 0 && b < 0 {
		fmt.Println("无法使用++操作实现")
		return -1
	}
	if b >= 0 {
		for i := 0; i < b; i++ {
			a++
		}
		return a
	} else {
		for i := 0; i < a; i++ {
			b++
		}
		return b
	}
}

//如何只使用++操作符实现加减乘除运算
func main() {
	fmt.Println("加法：")
	fmt.Println(add(2, 3))
}
