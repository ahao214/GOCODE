package main

import (
	"fmt"
	"math"
)

//获取最好的矩阵链相乘方法

//递归方法
func bestMatrixChainOrder(p []int, i, j int) int {
	if i == j {
		return 0
	}

	min := math.MaxInt64
	//通过把括号放在第一个不同的地方来获取最小的代价
	//每个括号内可以递归地使用相同的方法来计算
	for k := i; k < j; k++ {
		count := bestMatrixChainOrder(p, i, k) + bestMatrixChainOrder(p, k+1, j) + p[i-1]*p[k]*p[j]
		if count < min {
			min = count
		}
	}
	return min
}

func main() {
	arr := []int{1, 5, 2, 4, 6}
	fmt.Println("递归方法：")
	fmt.Println("最少的乘法次数为：", bestMatrixChainOrder(arr, 1, len(arr)-1))
}
