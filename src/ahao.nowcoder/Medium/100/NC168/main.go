package main

//NC168 盛水最多的容器
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, min(height[right], height[left])*(right-left))
		if height[right] > height[left] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
