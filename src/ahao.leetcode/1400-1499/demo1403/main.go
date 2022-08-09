package main

/*
1403. 非递增顺序的最小子序列
*/

//贪心
func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	tot := 0
	for _, num := range nums {
		tot += num
	}
	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		if sum > tot-sum {
			return nums[:i+1]
		}
	}
}
