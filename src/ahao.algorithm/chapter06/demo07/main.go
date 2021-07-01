package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//绝对值法
func max1(a, b int) int {
	if Abs(a-b) == a-b {
		return a
	} else {
		return b
	}
}

//二进制方法
func max2(a, b int) int {
	if (a-b)&(1<<31) != 0 {
		return b
	} else {
		return a
	}
}

//按要求比较两个数的大小
func main() {
	fmt.Println("绝对值方法：")
	fmt.Println(max1(5, 6))

	fmt.Println("二进制方法：")
	fmt.Println(max2(45, 56))
}
