package main

//题目：867.转置矩阵
//说明：给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[1,4,7],[2,5,8],[3,6,9]]
//解释：

func transpose(matrix [][]int) [][]int {
	row, col, result := len(matrix), len(matrix[0]), make([][]int, len(matrix[0]))
	for i := range result {
		result[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
