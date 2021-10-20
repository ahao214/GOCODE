package main

import "math"

//453. 最小操作次数使数组元素相等
func minMoves(nums []int) int {
	minValue := math.MaxInt32
	for _, v := range nums {
		if minValue > v {
			minValue = v
		}
	}
	res := 0
	for _, v := range nums {
		res += v - minValue
	}
	return res
}
