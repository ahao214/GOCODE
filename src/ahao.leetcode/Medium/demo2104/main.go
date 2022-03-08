package main

/*
2104. 子数组范围和
给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。

返回 nums 中 所有 子数组范围的 和 。

子数组是数组中一个连续 非空 的元素序列。
*/
func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		minVal, maxVal := num, num
		for _, v := range nums[i+1:] {
			if v < minVal {
				minVal = v
			} else if v > maxVal {
				maxVal = v
			}
			ans += int64(maxVal - minVal)
		}
	}
	return
}
