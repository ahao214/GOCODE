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

//按要求比较两个数的大小
func main() {
	fmt.Println("绝对值方法：")
	fmt.Println(max1(5, 6))
}
