package main

//1394. 找出数组中的幸运数
func findLucky(arr []int) int {
	res := -1
	if arr == nil || len(arr) == 0 {
		return res
	}
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	for k, v := range dic {
		if k == v && v > res {
			res = v
		}
	}
	return res
}
