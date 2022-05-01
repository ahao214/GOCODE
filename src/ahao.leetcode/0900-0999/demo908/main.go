package main

/*
908. 最小差值 I
*/

//数学
func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	ans := maxNum - minNum - 2*k
	if ans > 0 {
		return ans
	}
	return 0
}
