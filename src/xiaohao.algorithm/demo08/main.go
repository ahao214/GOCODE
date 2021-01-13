package main

import (
	"fmt"
)

//1.两数之和
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 3, 7, 9, 1, 4}
	target := 9
	fmt.Println(twoSum(nums, target))
}
