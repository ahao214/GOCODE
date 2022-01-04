package main

//51. N 皇后
func solveNQueens(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
	}

	//初始化棋盘
	board := [][]string{}
	for i := 0; i < n; i++ {
		row := []string{}
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	backtrack(board, 0, &res)
	return res
}

func backtrack(board [][]string, row int, res *[][]string) {
	if row == len(board) {
		*res = append(*res, covert(board))
		return
	}
	for col := 0; col < len(board[row]); col++ {
		if isValid(board, row, col) {
			//增加
			board[row][col] = "Q"
			//回溯
			backtrack(board, row+1, res)
			//移除
			board[row][col] = "."
		}
	}
}

func isValid(board [][]string, row int, col int) bool {
	//检查列
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == "Q" {
			return false
		}
	}
	m, n := row-1, col+1
	for m >= 0 && n < len(board) {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n++
	}
	m, n = row-1, col-1
	for m >= 0 && n >= 0 {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n--
	}
	return true
}

func covert(board [][]string) []string {
	res := []string{}
	for i := 0; i < len(board); i++ {
		row := ""
		for j := 0; j < len(board[i]); j++ {
			row += board[i][j]
		}
		res = append(res, row)
	}
	return res
}
