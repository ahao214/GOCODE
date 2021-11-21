package main

//NC177 打家劫舍(二)
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
