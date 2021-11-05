package main

import "sort"

//268. 丢失的数字
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}

//排序方法
func missingNumber2(arr []int) int {
	sort.Ints(arr)
	for i, n := range arr {
		if n != i {
			return i
		}
	}
	return len(arr)
}
