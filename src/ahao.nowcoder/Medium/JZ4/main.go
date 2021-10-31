package main

/**
 * JZ4 二维数组中的查找
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find(target int, array [][]int) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if array[i][j] == target {
				return true
			}
		}
	}
	return false
}
