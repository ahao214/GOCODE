package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//递归法
func getMinPathSub(arr [][]int, i, j int) int {
	//倒着走到了第一个结点，递归结束
	if i == 0 && j == 0 {
		return arr[i][j]
	} else if i > 0 && j > 0 {
		//选取两条可能路径上的最小值
		return arr[i][j] + Min(getMinPathSub(arr, i-1, j), getMinPathSub(arr, i, j-1))
	} else if i > 0 && j == 0 {
		return arr[i][j] + getMinPathSub(arr, i-1, j)
	} else {
		return arr[i][j] + getMinPathSub(arr, i, j-1)
	}
}

func getMinPath1(arr [][]int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	return getMinPathSub(arr, len(arr)-1, len(arr[0])-1)
}

//动态规划法
func getMinPath2(arr [][]int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	row, col := len(arr), len(arr[0])
	//用来保存计算的中间值
	cache := make([][]int, row)
	for i := range cache {
		cache[i] = make([]int, col)
	}
	cache[0][0] = arr[0][0]
	for i := 1; i < col; i++ {
		cache[0][i] = cache[0][i-1] + arr[0][i]
	}

	for j := 1; j < row; j++ {
		cache[j][0] = cache[j-1][0] + arr[j][0]
	}

	//在遍历二维数组的过程中不断把计算结果保存到cache中
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			//可以确定选择的路线为arr[i][j-1]
			if cache[i-1][j] > cache[i][j-1] {
				cache[i][j] = cache[i][j-1] + arr[i][j]
				fmt.Print("[", i, ",", j-1, "]  ")
			} else { //可以确定选择的路线为arr[i-1][j]
				cache[i][j] = cache[i-1][j] + arr[i][j]
				fmt.Print("[", i-1, ",", j, "]   ")
			}
		}
	}
	fmt.Print("[", row-1, ",", col-1, "]")
	fmt.Println()
	return cache[row-1][col-1]
}

//在二维数组中寻找最短路线
func main() {
	arr := [][]int{
		{1, 4, 3},
		{8, 7, 5},
		{2, 1, 5},
	}
	fmt.Println("递归法：")
	fmt.Println(getMinPath1(arr))

	fmt.Println("动态规划法：")
	fmt.Println(getMinPath2(arr))
}
