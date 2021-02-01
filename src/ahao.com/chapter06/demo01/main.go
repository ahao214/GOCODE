package main

import (
	"fmt"
)

//访问数组元素
func main() {
	arr := [5]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}
}
