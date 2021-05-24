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

//动态规划
func bestMatrixChainOrder2(p []int, n int) int {
	//申请数组来保存中间结果，为了简单不使用m[0][0]
	//cost[i,j]=计算A[i]*A[i+1]...*A[j]
	//所需的标量乘法的最小数量
	//其中A[i]的维数是p[i-1] x p[i]* /
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}
	//Len表示矩阵链的长度
	for cLen := 2; cLen < n; cLen++ {
		for i := 1; i < n-cLen+1; i++ {
			j := i + cLen - 1
			cost[i][j] = math.MaxInt64
			for k := i; k <= j-1; k++ {
				//计算乘法运算的代价
				q := cost[i][k] + cost[k+1][j] + p[i-1]*p[k]*p[j]
				if q < cost[i][j] {
					cost[i][j] = q
				}
			}
		}
	}
	return cost[1][n-1]
}

func main() {
	arr := []int{1, 5, 2, 4, 6}
	fmt.Println("递归方法：")
	fmt.Println("最少的乘法次数为：", bestMatrixChainOrder(arr, 1, len(arr)-1))

	fmt.Println("动态规划：")
	fmt.Println("最少的乘法次数为：", bestMatrixChainOrder2(arr, len(arr)))
}
