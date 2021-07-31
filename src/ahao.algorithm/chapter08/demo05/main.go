package main

import (
	"fmt"
)

func sort(array []int, low, high int) {
	if low >= high {
		return
	}
	i := low
	j := high
	var index int
	index = array[i]
	for i < j {
		for i < j && array[j] >= index {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < index {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = index
	sort(array, low, i-1)
	sort(array, i+1, high)
}

func QuickSort(array []int) {
	sort(array, 0, len(array)-1)
}

//快速排序
func main() {
	data := []int{5, 4, 9, 8, 1, 2, 3, 7, 6}
	QuickSort(data)
	fmt.Println(data)
}
