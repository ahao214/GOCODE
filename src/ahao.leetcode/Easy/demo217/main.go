package main

//217. 存在重复元素
func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	dic := make(map[int]int)
	for _, v := range nums {
		dic[v]++
	}

	for _, v := range dic {
		if v > 1 {
			return true
		}
	}
	return false
}
