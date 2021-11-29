package main

import "sort"

//786. 第 K 个最小的素数分数
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	type pair struct {
		x, y int
	}
	lst := make([]pair, 0, n*(n-1)/2)
	for i, x := range arr {
		for _, y := range arr[i+1:] {
			lst = append(lst, pair{x, y})
		}
	}
	sort.Slice(lst, func(i, j int) bool {
		a, b := lst[i], lst[j]
		return a.x*b.y < a.y*b.x
	})
	return []int{lst[k-1].x, lst[k-1].y}
}
