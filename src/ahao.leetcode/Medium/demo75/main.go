package main

//75. 颜色分类
func sortColors(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	left, right, i := 0, len(nums)-1, 0
	for i <= right && left < right {
		if nums[i] == 0 {
			nums[i] = nums[left]
			nums[left] = 0
			i++
			left++
		} else if nums[i] == 2 {
			nums[i] = nums[right]
			nums[right] = 2
			right--
		} else {
			i++
		}
	}
}
