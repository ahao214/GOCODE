package main

import "sort"

/**
 * 479 · 数组第二大数
 * @param nums: An integer array
 * @return: The second max number in the array.
 */
func secondMax(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-2]
}
