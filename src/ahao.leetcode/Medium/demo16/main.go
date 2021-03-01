package main

import (
	"fmt"
	"math"
	"sort"
)

//题目：16. 最接近的三数之和
//说明：给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)

func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
