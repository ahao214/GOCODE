package main

//566. 重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	if rows*cols != r*c {
		return mat
	} else {
		return reshape(mat, r, c)
	}
	return mat
}

func reshape(mat [][]int, r, c int) [][]int {
	newShape := make([][]int, r)
	for index := range newShape {
		newShape[index] = make([]int, c)
	}
	rowIndex, colIndex := 0, 0
	for _, row := range mat {
		for _, col := range row {
			if colIndex == c {
				colIndex = 0
				rowIndex++
			}
			newShape[rowIndex][colIndex] = col
			colIndex++
		}
	}
	return newShape
}
