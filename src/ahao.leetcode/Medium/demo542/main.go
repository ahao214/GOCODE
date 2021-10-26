package main

import "math"

//542. 01 矩阵
func updateMatrix(mat [][]int) [][]int {
	result := [][]int{}
	if len(mat) == 0 || len(mat[0]) == 0 {
		return result
	}
	maxRow, maxCol := len(mat), len(mat[0])
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if mat[r][c] == 1 && hasZero(mat, r, c) == false {
				// 将四周没有 0 的 1 特殊处理为最大值
				mat[r][c] = math.MaxInt64
			}
		}
	}
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if mat[r][c] == 1 {
				dfsMatrix(mat, r, c, -1)
			}
		}
	}
	return (mat)
}

//判断一个坐标的四周是否有零
func hasZero(matrix [][]int, row, col int) bool {
	if row > 0 && matrix[row-1][col] == 0 {
		return true
	}
	if col > 0 && matrix[row][col-1] == 0 {
		return true
	}
	if row < len(matrix)-1 && matrix[row+1][col] == 0 {
		return true
	}
	if col < len(matrix[0])-1 && matrix[row][col+1] == 0 {
		return true
	}
	return false
}

func dfsMatrix(matrix [][]int, row, col, val int) {
	// 不超过棋盘氛围，且 val 要比 matrix[row][col] 小
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || (matrix[row][col] <= val) {
		return
	}
	if val > 0 {
		matrix[row][col] = val
	}
	dfsMatrix(matrix, row-1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col-1, matrix[row][col]+1)
	dfsMatrix(matrix, row+1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col+1, matrix[row][col]+1)
}
