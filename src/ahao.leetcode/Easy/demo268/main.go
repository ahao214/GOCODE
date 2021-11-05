package main

import "sort"

//268. 丢失的数字

//位运算
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

//哈希集合
func missingNumber3(arr []int) int {
	has := map[int]bool{}
	for _, v := range arr {
		has[v] = true
	}
	for i := 0; ; i++ {
		if !has[i] {
			return i
		}
	}
}
