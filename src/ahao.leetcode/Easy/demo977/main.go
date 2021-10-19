package main

//977. 有序数组的平方
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left := 0
	right := n - 1
	for p := n - 1; p >= 0; p-- {
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]
		if leftVal > rightVal {
			res[p] = leftVal
			left++
		} else {
			res[p] = rightVal
			right--
		}
	}
	return res
}
