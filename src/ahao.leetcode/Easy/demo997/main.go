package main

//997. 找到小镇的法官
func findJudge(n int, trust [][]int) int {
	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)
	for _, t := range trust {
		inDegrees[t[1]]++
		outDegrees[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}
	return -1
}

func findJudge1(n int, trust [][]int) int {
	if n == 1 {
		return n
	}

	dic := make(map[int]int)
	for _, v := range trust {
		dic[v[1]]++
	}
	res := -1
	for k, v := range dic {
		if v == n-1 {
			res = k
			break
		}
	}
	if res > -1 {
		for _, v := range trust {
			if v[0] == res {
				return -1
			}
		}
	}
	return res
}
