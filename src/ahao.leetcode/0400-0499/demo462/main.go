package main

import (
	"math/rand"
	"sort"
	"time"
)

/*
462. 最少移动次数使数组元素相等 II
给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。

在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。



示例 1：

输入：nums = [1,2,3]
输出：2
解释：
只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
示例 2：

输入：nums = [1,10,2,9]
输出：16


提示：

n == nums.length
1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

//排序
func minMoves2(nums []int) (ans int) {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//快速选择
func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func randomPartition(a []int, l, r int) int {
	i := rand.Intn(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	}
	if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func minMoves22(nums []int) (ans int) {
	rand.Seed(Now().UnixNano())
	x := quickSelect(nums, 0, len(nums)-1, len(nums)/2)
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}
