package main

import (
	"fmt"
)

func Merge(array []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1)
	R := make([]int, n2)
	for i, k := 0, p; i < n1; i, k = i+1, k+1 {
		L[i] = array[k]
	}
	for i, k := 0, q+1; i < n2; i, k = i+1, k+1 {
		R[i] = array[k]
	}
	var i, k, j int
	for i, k, j = 0, p, 0; i < n1 && j < n2; k++ {
		if L[i] < R[j] {
			array[k] = L[i]
			i++
		} else {
			array[k] = R[j]
			j++
		}
		if i < n1 {
			for j = i; j < n1; j, k = j+1, k+1 {
				array[k] = L[j]
			}
		}
		if j < n2 {
			for i = j; i < n2; i, k = i+1, k+1 {
				array[k] = R[i]
			}
		}
	}
}

func MergeSort(array []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(array, p, q)
		MergeSort(array, q+1, r)
		Merge(array, p, q, r)
	}
}

//归并排序
func main() {
	array := []int{3, 2, 1, 4, 6, 8, 9}
	MergeSort(array, 0, len(array)-1)
	fmt.Println(array)
}
