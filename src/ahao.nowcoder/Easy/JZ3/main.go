package main

/**
 * JZ3 数组中重复的数字
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	if numbers == nil || len(numbers) == 0 {
		return -1
	}
	dic := make(map[int]int)
	for _, v := range numbers {
		dic[v]++
	}
	for k, v := range dic {
		if v > 1 {
			return k
		}
	}
	return -1
}
