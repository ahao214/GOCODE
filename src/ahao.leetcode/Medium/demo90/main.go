package main

import (
	"sort"
)

//90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}
	visited := make([]bool, len(nums))
	cur := []int{}
	sort.Ints(nums)
	backtrack(nums, 0, cur, visited, &res)
	return res
}

func backtrack(arr []int, start int, cur []int, visited []bool, res *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	for i := start; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] && visited[i-1] == false {
			continue
		}
		cur = append(cur, arr[i])
		visited[i] = true
		backtrack(arr, i+1, cur, visited, res)
		visited[i] = false
		//移除
		cur = cur[:len(cur)-1]
	}
}
