package main

import (
	"strconv"
)

//60. 排列序列
/*
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。
*/
func getPermutation(n int, k int) string {
	res := 0
	cur := 0
	count := 0
	visited := make([]bool, n)
	backtrack(n, k, visited, cur, &count, &res)
	return strconv.Itoa(res)
}

func backtrack(n int, k int, visited []bool, cur int, count *int, res *int) {
	if lengthOfNumber(cur) == n {
		(*count)++
		if (*count) == k {
			(*res) = cur
			return
		}
	}
	for i := 1; i <= n; i++ {
		if visited[i-1] {
			continue
		}
		cur = cur*10 + i
		visited[i-1] = true
		backtrack(n, k, visited, cur, count, res)
		visited[i-1] = false
		cur = (cur - i) / 10
	}
}

func lengthOfNumber(n int) int {
	res := 0
	for n >= 1 {
		n = n / 10
		res++
	}
	return res
}
