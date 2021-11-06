package main

import "fmt"

/**
 * 连续子数组的最大和
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	sum := 0
	result := array[0]
	for _, v := range array {
		sum = max(sum, 0) + v
		result = max(sum, result)
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1}
	fmt.Println(FindGreatestSumOfSubArray(arr))
}
