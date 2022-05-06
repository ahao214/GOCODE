package main

import (
	"math"
	"sort"
)

/*
713. 乘积小于 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
*/

//二分查找
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k == 0 {
		return
	}
	n := len(nums)
	logPrefix := make([]float64, n+1)
	for i, num := range nums {
		logPrefix[i+1] = logPrefix[i] + math.Log(float64(num))
	}
	logK := math.Log(float64(k))
	for j := 1; j <= n; j++ {
		l := sort.SearchFloat64s(logPrefix[:j], logPrefix[j]-logK+1e-10)
		ans += j - l
	}
	return
}

//滑动窗口
func numSubarrayProductLessThanK2(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}
