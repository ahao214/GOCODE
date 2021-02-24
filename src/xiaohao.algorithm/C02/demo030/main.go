package main

import "fmt"

//第136题：只出现一次的数字
func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 2, 3, 4, 4}
	fmt.Println(singleNumber(nums))
}
