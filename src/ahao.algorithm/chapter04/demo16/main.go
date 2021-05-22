package main

import (
	"fmt"
)

//在有规律的二维数组中进行高效的数据查找
func findWithBinary(arr [][]int, data int) bool {
	if arr == nil {
		return false
	}

	//从二维数组右上角元素开始遍历
	rows := len(arr)
	columns := len(arr[0])
	for i, j := 0, columns-1; i < rows && j >= 0; {
		//在数组中找到data，返回
		if arr[i][j] == data {
			return true
		} else if arr[i][j] > data {
			//当前遍历到数组中的值大于data，data肯定不在这一列中
			j--
		} else {
			//当前遍历到数组中的值小于data，data肯定不在这一行中
			i++
		}
	}
	return false
}

func main() {
	arr := [][]int{
		{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34},
		{40, 41, 42, 43, 44},
	}

	fmt.Println(findWithBinary(arr, 17))
	fmt.Println(findWithBinary(arr, 14))

}
