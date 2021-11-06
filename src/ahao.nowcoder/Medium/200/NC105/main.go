package main

import (
	"fmt"
)

/**
 * NC105 二分查找-II
 *
 * 如果目标值存在返回下标，否则返回 -1
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search(nums []int, target int) int {
	res := -1
	if nums == nil || len(nums) == 0 {
		return res
	}
	if nums[0] == target {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			if nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func main() {
	arr := []int{1, 1, 2, 3, 7, 7, 7, 9, 9, 10}
	k := 2
	res := search(arr, k)
	fmt.Println(res)
}
