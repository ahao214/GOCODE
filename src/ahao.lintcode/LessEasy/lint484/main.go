package main

/**
 * 484 · 交换数组两个元素
 * @param A: An integer array
 * @param index1: the first index
 * @param index2: the second index
 * @return: nothing
 */
func swapIntegers(A *[]int, index1 int, index2 int) {
	temp := (*A)[index1]
	(*A)[index1] = (*A)[index2]
	(*A)[index2] = temp
}
