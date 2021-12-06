package main

//343. 整数拆分
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak2(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
