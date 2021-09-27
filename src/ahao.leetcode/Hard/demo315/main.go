package main

import (
	"github.com/halfrost/LeetCode-Go/template"
	"sort"
)

//315. 计算右侧小于当前元素的个数
func countSmaller1(nums []int) []int {
	// copy 一份原数组至所有数字 allNums 数组中
	allNums, res := make([]int, len(nums)), []int{}
	copy(allNums, nums)
	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}
	// 树状数组 Query
	bit := template.BinaryIndexedTree{}
	bit.Init(k)
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.Query(kth[nums[i]]-1))
		bit.Add(kth[nums[i]], 1)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
