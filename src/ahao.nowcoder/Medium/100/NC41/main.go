package main

/**
 * NC41 最长无重复子数组
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxLength(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	res := 0
	left := -1
	dic := make(map[int]int)
	for right := 0; right < len(arr); right++ {
		if _, ok := dic[arr[right]]; ok {
			left = max(left, dic[arr[right]])
		}
		res = max(res, right-left)
		dic[arr[right]] = right
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
