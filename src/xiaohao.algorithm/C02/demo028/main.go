package main

import "fmt"

//905.按奇偶排序数组
func sortArrayByParity(A []int) []int {
	j := 0
	for i := range A {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}
	return A
}

func main() {
	nums := []int{2, 3, 1, 4}
	arr := sortArrayByParity(nums)
	fmt.Println(arr)
}
