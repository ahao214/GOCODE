package main

import "fmt"

func adjustMinHeap(a []int, pos, llen int) {
	var temp int
	child := 0

	for temp = a[pos]; 2*pos+1 <= llen; pos = child {
		child = 2*pos + 1
		if child < llen && a[child] > a[child+1] {
			child++
		}
		if a[child] < temp {
			a[pos] = a[child]
		} else {
			break
		}
	}
	a[pos] = temp
}

func MinHeapSort(array []int) {
	llen := len(array)
	for i := llen/2 - 1; i >= 0; i-- {
		adjustMinHeap(array, i, llen-1)
	}
	for i := llen - 1; i >= 0; i-- {
		tmp := array[0]
		array[0] = array[i]
		array[i] = tmp
		adjustMinHeap(array, 0, i-1)
	}
}

//堆排序
func main() {
	fmt.Println("堆排序：")
	data := []int{4, 5, 6, 7, 8, 9, 0, 3, 1, 2}
	MinHeapSort(data)
	fmt.Println(data)
}
