package main

/**
 * 767.翻转数组
 * @param nums: a integer array
 * @return: nothing
 */
func reverseArray(nums *[]int) {
	left := 0
	right := len(*nums) - 1
	for left <= right {
		tmp := (*nums)[left]
		(*nums)[left] = (*nums)[right]
		(*nums)[right] = tmp
		left++
		right--
	}
}
