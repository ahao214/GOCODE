package main

/**
 * 1334 · 旋转数组
 * @param nums: an array
 * @param k: an integer
 * @return: rotate the array to the right by k steps
 */
func rotate(nums []int, k int) []int {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
	return nums
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
