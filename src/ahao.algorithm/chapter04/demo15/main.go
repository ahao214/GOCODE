package main

import (
	"fmt"
)

//对数组进行循环移位

//蛮力法
func rightShift1(arr []int, k int) {
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

func rightShift2(arr []int, k int) {
	if arr == nil {
		fmt.Println("参数不合法")
		return
	}
	ll := len(arr)
	k %= ll
	for k != 0 {
		k--
		t := arr[ll-1]
		for i := ll - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = t
	}
}

//翻转法
func reverse(arr []int, start, end int) {
	for start < end {
		tmp := arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
		start++
		end--
	}
}

func rightShift3(arr []int, k int) {
	if arr == nil {
		fmt.Println("参数不合法")
		return
	}
	ll := len(arr)
	k %= ll
	reverse(arr, 0, ll-k-1)
	reverse(arr, ll-k, ll-1)
	reverse(arr, 0, ll-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rightShift1(arr, 4)
	fmt.Println("蛮力法：")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println("--")
	rightShift2(arr, 4)
	for _, v := range arr {
		fmt.Print(v, " ")
	}

	fmt.Println("翻转法")
	rightShift3(arr, 4)
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
