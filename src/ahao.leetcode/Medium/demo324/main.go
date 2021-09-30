package main

import "sort"

//324. 摆动排序 II
func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	array := make([]int, len(nums))
	copy(array, nums)
	sort.Ints(array)
	n := len(nums)
	left := (n+1)/2 - 1 // median index
	right := n - 1      // largest value index
	for i := 0; i < len(nums); i++ {
		// copy large values on odd indexes
		if i%2 == 1 {
			nums[i] = array[right]
			right--
		} else { // copy values decremeting from median on even indexes
			nums[i] = array[left]
			left--
		}
	}
}
