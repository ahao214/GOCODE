package main

//27. 移除元素
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res += 1
		}
	}
	return res
}
