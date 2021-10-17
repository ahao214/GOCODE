package main

/**
 * 463 · 整数排序
 * @param A: an integer array
 * @return: nothing
 */
func sortIntegers(A *[]int) {
	for i := 0; i < len(*A)-1; i++ {
		for j := len(*A) - 1; j > i; j-- {
			if (*A)[j] < (*A)[j-1] {
				tmp := (*A)[j]
				(*A)[j] = (*A)[j-1]
				(*A)[j-1] = tmp
			}
		}
	}
}
