package main

import "sort"

//594. 最长和谐子序列

//枚举
func findLHS(nums []int) int {
	ans := 0
	sort.Ints(nums)
	begin := 0
	for end, n := range nums {
		if n-nums[begin] > 1 {
			begin++
		}
		if n-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return ans
}
