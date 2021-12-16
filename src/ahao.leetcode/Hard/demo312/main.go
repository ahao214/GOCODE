package main

//312. 戳气球
func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	return solve(0, n+1, val, rec)
}

func solve(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += solve(left, i, val, rec) + solve(i, right, val, rec)
		rec[left][right] = max(rec[left][right], sum)
	}
	return rec[left][right]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//动态规划
func maxCoins1(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}
