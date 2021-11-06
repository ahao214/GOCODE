package main

/**
 * 求平方根-向下取整
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 1
	right := x / 2
	for left < right {
		mid := left + (right-left+1)>>1
		temp := x / mid
		if mid > temp {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
