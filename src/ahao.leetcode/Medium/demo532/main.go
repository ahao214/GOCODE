package main

//532. 数组中的 k-diff 数对
func findPairs(nums []int, k int) int {
	if nums == nil || len(nums) < 2 || k < 0 {
		return 0
	}
	res := 0
	dic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tmp := nums[i] + k
		index, ok := dic[tmp]
		if ok && i != index {
			res++
			delete(dic, tmp)
		}
	}
	return res
}
