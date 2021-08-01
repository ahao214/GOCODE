package main

import "fmt"

func shellSort(array []int) {
	llen := len(array)
	for h := llen / 2; h > 0; h = h / 2 {
		for i := h; i < llen; i++ {
			temp := array[i]
			var j int
			for j = i - h; j >= 0; j -= h {
				if temp < array[j] {
					array[j+h] = array[j]
				} else {
					break
				}
			}
			array[j+h] = temp
		}
	}
}

//希尔排序
func main() {
	data := []int{3, 4, 2, 1, 5, 7, 8, 9, 6}
	shellSort(data)
	fmt.Println(data)
}
