package main

/**
 * NC140 排序
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	llen := len(arr)
	for i := 0; i < llen-1; i++ {
		for j := llen - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}
	return arr
}
