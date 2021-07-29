package main

import (
	"fmt"
)

func BubbleSort(array []int) {
	len := len(array)
	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			if array[j] < array[j-1] {
				tmp := array[j]
				array[j] = array[j-1]
				array[j-1] = tmp
			}
		}
	}
}

//冒泡排序
func main() {
	array := []int{4, 5, 1, 2, 3, 6, 8, 9, 7}
	BubbleSort(array)
	fmt.Println(array)
}
