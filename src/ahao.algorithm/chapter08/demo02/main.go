package main

import (
	"fmt"
)

func insertSort(array []int) {
	if array == nil {
		return
	}
	for i := 1; i < len(array); i++ {
		tmp, j := array[i], i
		if array[j-1] > tmp {
			for j >= 1 && array[j-1] > tmp {
				array[j] = array[j-1]
				j--
			}
		}
		array[j] = tmp
	}
}

//插入排序
func main() {
	array := []int{2, 3, 1, 4, 5, 9, 7, 8, 6}
	insertSort(array)
	fmt.Println(array)
}
