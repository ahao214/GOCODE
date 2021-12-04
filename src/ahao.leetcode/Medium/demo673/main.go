package main

//673. 最长递增子序列的个数
func findNumberOfLIS(nums []int) int {
	n, max := len(nums), -1
	dp, count := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i], count[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] += count[j]
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if max == dp[i] {
			ans += count[i]
		}
	}
	return ans
}
