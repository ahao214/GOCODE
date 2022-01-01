package main

//39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return res
	}
	cur := []int{}
	backtrack(candidates, target, 0, cur, &res)
	return res
}

func backtrack(arr []int, target int, start int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(arr); i++ {
		cur = append(cur, arr[i])
		backtrack(arr, target-arr[i], i, cur, res)
		cur = cur[:len(cur)-1]
	}
}
