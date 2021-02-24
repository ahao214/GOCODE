package main

import "fmt"

//插入排序
func main() {
	arr := []int{5, 4, 3, 2, 1}
	insertSort(arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		fmt.Println(arr)
	}
}
