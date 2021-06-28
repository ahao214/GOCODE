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

//用++实现减法操作(限制条件：被减数大于减数)
func minus(a, b int) int {
	if a < b {
		fmt.Println("无法使用++实现减法操作")
		return -1
	}
	result := 0
	for ; b != a; b++ {
		result++
	}
	return result
}

//用++实现乘法操作(限制条件：两个数都为整数)
func multi(a, b int) int {
	if a <= 0 || b <= 0 {
		fmt.Println("无法使用++实现乘法操作")
		return -1
	}
	res := 0
	for i := 0; i < b; i++ {
		res = add(res, a)
	}
	return res
}

//如何只使用++操作符实现加减乘除运算
func main() {
	fmt.Println("加法：")
	fmt.Println(add(2, 3))
	fmt.Println("减法：")
	fmt.Println(minus(9, 2))
	fmt.Println("乘法：")
	fmt.Println(multi(9, 3))
}
