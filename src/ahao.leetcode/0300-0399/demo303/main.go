package main

//题目：303. 区域和检索 - 数组不可变
//说明：给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点
//输入：["NumArray", "sumRange", "sumRange", "sumRange"]
//[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
//输出：[null, 1, -1, -3]
//解释：NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
//numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
//numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
//numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{prefixSum: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.prefixSum[j] - this.prefixSum[i-1]
	}
	return this.prefixSum[j]
}
