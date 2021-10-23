package main

//73.矩阵置零
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	//行数
	rows := len(matrix)
	//列数
	cols := len(matrix[0])
	i := 0
	j := 0
	for i < rows && j < cols {
		if matrix[i][j] == 0 {
			setZero(i, j, matrix)
			i++
			j++
		}
	}
}

func setZero(r, c int, matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][c] = 0
	}
	for i := 0; i < len(matrix[0]); i++ {
		matrix[r][i] = 0
	}
}

func setZeroesT(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	isFirstRowExistZero, isFirstColExistZero := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isFirstColExistZero = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			isFirstRowExistZero = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 处理[1:]行全部置 0
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	// 处理[1:]列全部置 0
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if isFirstRowExistZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if isFirstColExistZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
