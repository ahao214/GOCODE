package main

//851. 喧闹和富有

//深度优先
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		if ans[x] != -1 {
			return
		}
		ans[x] = x
		for _, y := range g[x] {
			dfs(y)
			if quiet[ans[y]] < quiet[ans[x]] {
				ans[x] = ans[y]
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return ans
}

//拓扑排序
func loudAndRich1(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	inDeg := make([]int, n)
	for _, r := range richer {
		g[r[0]] = append(g[r[0]], r[1])
		inDeg[r[1]]++
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	q := make([]int, 0, n)
	for i, deg := range inDeg {
		if deg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			if quiet[ans[x]] < quiet[ans[y]] {
				ans[y] = ans[x] // 更新 x 的邻居的答案
			}
			inDeg[y]--
			if inDeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return ans
}
