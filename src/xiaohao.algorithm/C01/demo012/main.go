package main

//120.三角形最小路径和
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle))
	for i, arr := range triangle {
		dp[i] = make([]int, len(arr))
	}
	result := 1<<31 - 1
	dp[0][0] = triangle[0][0]
	dp[1][1] = triangle[1][1] + triangle[0][0]
	dp[1][0] = triangle[1][0] + triangle[0][0]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	for _, k := range dp[len(dp)-1] {
		result = min(result, k)
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}
