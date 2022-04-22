package main

/*
821. 字符的最短距离
给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。

返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。

两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。
*/
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}

	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
