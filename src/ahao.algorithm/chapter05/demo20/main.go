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

//在二维数组中寻找最短路线
func main() {
	arr := [][]int{
		{1, 4, 3},
		{8, 7, 5},
		{2, 1, 5},
	}
	fmt.Println("递归法：")
	fmt.Println(getMinPath1(arr))
}
