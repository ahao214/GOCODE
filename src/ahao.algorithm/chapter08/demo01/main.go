package main

import (
	"fmt"
)

func selectSort(data []int) {
	len := len(data)
	for i := 0; i < len; i++ {
		tmp := data[i]
		flag := i
		for j := i + 1; j < len; j++ {
			if data[j] < tmp {
				tmp = data[j]
				flag = j
			}
		}
		if flag != i {
			data[flag] = data[i]
			data[i] = tmp
		}
	}
}

//选择排序
func main() {
	data := []int{4, 5, 1, 2, 3, 6, 9, 8, 7}
	selectSort(data)
	fmt.Println(data)
}
