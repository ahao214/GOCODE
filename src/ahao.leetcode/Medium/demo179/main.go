package main

import (
	"sort"
	"strconv"
)

/*
179. 最大数
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/
func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}
	var number string
	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		if first+second >= second+first {
			return true
		}
		return false
	})
	for _, num := range nums {
		number += strconv.Itoa(num)
	}
	if number[0] == '0' {
		return number[:1]
	}
	return number
}
