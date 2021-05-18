package main

import (
	"fmt"
	"math"
)

//求数组中连续最大和

//蛮力法
func maxSubArray(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}

	MaxSum := 0
	len := len(arr)
	i := 0
	j := 0
	k := 0
	for i = 0; i < len; i++ {
		for j = i; j < len; j++ {
			ThisSum := 0
			for k = i; k < j; k++ {
				ThisSum += arr[k]
			}
			if ThisSum > MaxSum {
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

//重复利用已经计算的子数组和
func maxSubArraySum(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	len := len(arr)
	maxSum := math.MinInt64
	for i := 0; i < len; i++ {
		sum := 0
		for j := i; j < len; j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

//动态规范方法
func maxSubArrayDyn(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	n := len(arr)
	End := make([]int, n)
	All := make([]int, n)
	End[n-1] = arr[n-1]
	All[n-1] = arr[n-1]
	End[0] = arr[0]
	All[0] = arr[0]
	for i := 1; i < n; i++ {
		End[i] = Max(End[i-1]+arr[i], arr[i])
		All[i] = Max(End[i], All[i-1])
	}
	return All[n-1]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, -2, 4, 8, -4, 7, -1, -5}
	fmt.Println("蛮力法")
	fmt.Println("连续最大和为：", maxSubArray(arr))
	fmt.Println("重复利用已经计算的子数组和")
	fmt.Println("连续最大和为：", maxSubArraySum(arr))
	fmt.Println("动态规划法")
	fmt.Println("连续最大和为：", maxSubArrayDyn(arr))
}
