package main

import (
	"fmt"
)

//第268题：缺失数字
func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	res := 0
	for i, k := range nums {
		res ^= k ^ i
	}
	return res ^ len(nums)
}
