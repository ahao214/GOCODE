package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"math"
)

//求数组中绝对值最小的数

//顺序比较法
func FindMin(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	min := math.MaxInt64
	for _, v := range arr {
		if Abs(v) < Abs(min) {
			min = v
		}
	}
	return min
}

func main() {
	arr := []int{-10, -4, -2, 7, 15, 50}
	fmt.Println("顺序比较法：")
	fmt.Println("绝对值最小的数为：", FindMin(arr))
}
