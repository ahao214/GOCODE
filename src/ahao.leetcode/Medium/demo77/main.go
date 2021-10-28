package main

//77. 组合
func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 0 || n == 0 {
		return res
	}
	cur := []int{}
	backtrack(n, k, 1, cur, &res)
	return res
}

func backtrack(n int, k int, start int, cur []int, res *[][]int) {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		backtrack(n, k, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}
