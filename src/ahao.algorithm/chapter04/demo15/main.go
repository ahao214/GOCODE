package main

import (
	"fmt"
)

//对数组进行循环移位

//蛮力法
func rightShift(arr []int, k int) {
	if arr == nil {
		fmt.Println("参数不合法")
		return
	}
	ll := len(arr)
	for k != 0 {
		k--
		tmp := arr[ll-1]
		for i := ll - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = tmp
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rightShift(arr, 4)
	fmt.Println("蛮力法：")
	for _, v := range arr {
		fmt.Print(v, " ")
	}

}
