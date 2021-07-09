package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count*2 > len(nums) {
		return candidate
	}
	return -1
}

//面试题 17.10. 主要元素
func main() {
	array := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	fmt.Println(majorityElement(array))
}
