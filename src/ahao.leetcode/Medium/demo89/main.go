package main

//89. 格雷编码
func grayCode(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|1<<(i-1))
		}
	}
	return ans
}

func grayCode1(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = 1>>1 ^ i
	}
	return ans
}
