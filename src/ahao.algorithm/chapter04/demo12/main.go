package main

import (
	"fmt"
)

//旋转数组
func rotateArr(arr [][]int) {
	row := 0
	col := 0
	len := len(arr)
	//打印二维数组右上半部分
	for i := len - 1; i > 0; i-- {
		row = 0
		col = i
		for col < len {
			fmt.Print(arr[row][col], " ")
			row++
			col++
		}
		fmt.Println()
	}

	//打印二维数组左下半部分
	for i := 0; i < len; i++ {
		row = i
		col = 0
		for row < len {
			fmt.Print(arr[row][col], " ")
			row++
			col++
		}
		fmt.Println()
	}
}

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateArr(arr)
}
