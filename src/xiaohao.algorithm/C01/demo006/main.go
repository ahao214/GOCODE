package main

import (
	"fmt"
)

//26:删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{0, 1, 1, 2, 3, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(nums))
}
