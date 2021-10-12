package main

/**
 * 螺旋矩阵
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1

	var x, y int
	for left <= right && up <= down {
		for y = left; y <= right && up <= down && left <= right; y++ {
			result = append(result, matrix[x][y])
		}
		y--
		up++
		for x = up; x <= down && up <= down && left <= right; x++ {
			result = append(result, matrix[x][y])
		}
		x--
		right--
		for y = right; y >= left && up <= down && left <= right; y-- {
			result = append(result, matrix[x][y])
		}
		y++
		down--
		for x = down; x >= up && up <= down && left <= right; x-- {
			result = append(result, matrix[x][y])
		}
		x++
		left++
	}
	return result
}
