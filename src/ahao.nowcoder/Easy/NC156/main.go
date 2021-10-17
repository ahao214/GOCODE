package main

/**
 * 数组中只出现一次的数（其它数出现k次）
 * @param arr int一维数组
 * @param k int
 * @return int
 */
func foundOnceNumber(arr []int, k int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	for c, v := range dic {
		if v == 1 {
			return c
		}
	}
	return -1
}
