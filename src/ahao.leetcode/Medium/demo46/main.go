package main

//46. 全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 || nums == nil {
		return res
	}
	visited := make([]bool, len(nums))
	cur := []int{}
	backtrack(nums, cur, visited, &res)
	return res
}

func backtrack(arr []int, cur []int, visited []bool, res *[][]int) {
	if len(arr) == len(cur) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(arr); i++ {
		if visited[i] {
			continue
		}
		cur = append(cur, arr[i])
		visited[i] = true
		backtrack(arr, cur, visited, res)
		visited[i] = false
		cur = cur[:len(cur)-1]
	}
}
