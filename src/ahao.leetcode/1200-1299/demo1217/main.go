package main

/*
1217. 玩筹码
*/

//贪心
func minCostToMoveChips(position []int) int {
	cnt := [2]int{}
	for _, p := range position {
		cnt[p%2]++
	}
	return min(cnt[0], cnt[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
