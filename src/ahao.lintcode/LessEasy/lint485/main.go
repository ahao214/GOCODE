package main

/**
 * 485 · 生成给定大小的数组
 * @param size: An integer
 * @return: An integer list
 */
func generate(size int) []int {
	res := []int{}
	for i := 1; i <= size; i++ {
		res = append(res, i)
	}
	return res
}
