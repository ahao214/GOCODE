package main

import (
	"fmt"
)

//260. 只出现一次的数字 III
func singleNumber(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	diff &= -diff
	result := []int{0, 0}
	for _, num := range nums {
		if (num & diff) == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	result := singleNumber(nums)
	for _, value := range result {
		fmt.Println(value)
	}
}
