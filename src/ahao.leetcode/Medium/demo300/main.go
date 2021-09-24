package main

import (
	"github.com/isdamir/gotype"
	"sort"
)

//300. 最长递增子序列

//方法一：DP
func longestofLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = gotype.Max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = gotype.Max(res, dp[i])
	}
	return res
}

//方法二：
func longestLIS1(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}
