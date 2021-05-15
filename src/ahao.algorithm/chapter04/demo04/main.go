package main

import (
	"fmt"
)

//找出数组中丢失的数

//累加求和方法
func getNum(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	suma := 0
	sumb := 0
	for i, v := range arr {
		suma += v
		sumb += i
	}
	sumb = sumb + len(arr)*2 + 1
	return sumb - suma
}

//异或方法
func getNumXor(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	a := arr[0]
	b := 1
	len := len(arr)
	for i := 1; i < len; i++ {
		a ^= arr[i]
	}
	for i := 2; i <= len+1; i++ {
		b ^= i
	}
	return a ^ b
}

func main() {
	fmt.Println("累加求和：")
	arr := []int{1, 4, 3, 2, 7, 5}
	fmt.Println(getNum(arr))
	fmt.Println("异或方法")
	fmt.Println(getNumXor(arr))
}
