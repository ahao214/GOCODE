package main

/*
1394. 找出数组中的幸运数

在整数数组中，如果一个整数的出现频次和它的数值大小相等，我们就称这个整数为「幸运数」。
给你一个整数数组 arr，请你从中找出并返回一个幸运数。
如果数组中存在多个幸运数，只需返回 最大 的那个。
如果数组中不含幸运数，则返回 -1 。
*/
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
