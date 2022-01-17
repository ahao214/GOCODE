package main

/*
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	threshold := 1 + len(nums)/2
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val, ok := m[nums[i]]
		if ok {
			val++
			if val >= threshold {
				return nums[i]
			} else {
				m[nums[i]] = val
			}
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}
