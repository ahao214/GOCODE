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

func moveZeros(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	left := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left = i
			break
		}
	}
	if left < 0 {
		return
	}
	for i := left; i < len(nums); i++ {
		if nums[i] != 0 && nums[left] == 0 {
			tmp := nums[i]
			nums[i] = nums[left]
			nums[left] = tmp
			left++
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
