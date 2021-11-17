package main

//NC160 二分查找
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	res := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
