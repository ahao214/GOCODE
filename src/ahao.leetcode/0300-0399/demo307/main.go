package main

import "math"

/*
307. 区域和检索 - 数组可修改
给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
*/
type NumArray struct {
	nums, sums []int // sums[i] 表示第 i 个块的元素和
	size       int   // 块的大小
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	size := int(math.Sqrt(float64(n)))
	sums := make([]int, (n+size-1)/size) // n/size 向上取整
	for i, num := range nums {
		sums[i/size] += num
	}
	return NumArray{nums, sums, size}
}

func (na *NumArray) Update(index, val int) {
	na.sums[index/na.size] += val - na.nums[index]
	na.nums[index] = val
}

func (na *NumArray) SumRange(left, right int) (ans int) {
	size := na.size
	b1, b2 := left/size, right/size
	if b1 == b2 { // 区间 [left, right] 在同一块中
		for i := left; i <= right; i++ {
			ans += na.nums[i]
		}
		return
	}
	for i := left; i < (b1+1)*size; i++ {
		ans += na.nums[i]
	}
	for i := b1 + 1; i < b2; i++ {
		ans += na.sums[i]
	}
	for i := b2 * size; i <= right; i++ {
		ans += na.nums[i]
	}
	return
}
