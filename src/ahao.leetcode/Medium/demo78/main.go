package main

//78. 子集
func subsets(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}
	cur := []int{}
	backtrack(nums, 0, cur, &res)
	return res
}

func backtrack(arr []int, start int, cur []int, res *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	for i := start; i < len(arr); i++ {
		cur = append(cur, arr[i])
		backtrack(arr, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}
