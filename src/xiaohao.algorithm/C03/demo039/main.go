package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	sub := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
		sub = nums[i+1] - nums[i]
	}
	return sub < 5
}

//扑克牌中的顺子
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(isStraight(nums))
}
