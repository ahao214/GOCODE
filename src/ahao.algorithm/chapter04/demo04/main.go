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

func main() {
	fmt.Println("累加求和：")
	arr := []int{1, 4, 3, 2, 7, 5}
	fmt.Println(getNum(arr))
}
