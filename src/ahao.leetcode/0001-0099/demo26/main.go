package main

//26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func removeDuplicates2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[res] {
			res++
			nums[res] = nums[i]
		}
	}
	return res + 1
}
