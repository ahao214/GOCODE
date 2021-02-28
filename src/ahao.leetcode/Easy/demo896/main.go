package main

import (
	"fmt"
)

//题目：896. 单调数列
//说明：如果数组是单调递增或单调递减的，那么它是单调的。
//如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。
//如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。
//当给定的数组 A 是单调数组时返回 true，否则返回 false。

//输入：[1,2,2,3]
//输出：true

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	if A[0] < A[1] {
		return up(A[1:])
	}
	if A[0] > A[1] {
		return down(A[1:])
	}
	return up(A[1:]) || down(A[1:])
}

//判断数组是否是单调递增
func up(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

//判断数组是否是单调递减
func down(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 2, 3}
	fmt.Println(isMonotonic(arr))
}
