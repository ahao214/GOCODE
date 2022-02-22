package main

//NC128 接雨水问题
func maxWater(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	N := len(arr)
	left := 1
	leftMax := arr[0]
	right := N - 2
	rightMax := arr[N-1]
	water := 0
	for left <= right {
		if leftMax <= rightMax {
			water += max(0, leftMax-arr[left])
			leftMax = max(leftMax, arr[left])
			left++
		} else {
			water += max(0, rightMax-arr[right])
			rightMax = max(rightMax, arr[right])
			right--
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
