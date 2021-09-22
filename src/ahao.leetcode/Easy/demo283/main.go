package main

import (
	"fmt"
)

//283. 移动零
func moveZero(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
}

func main() {
	nums := []int{0, 1, 2, 0, 3}
	moveZero(nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}
