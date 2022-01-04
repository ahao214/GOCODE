package main

//52. N皇后 II
func totalNQueens(n int) int {
	res := 0
	if n == 0 {
		return res
	}
	board := [][]string{}
	for i := 0; i < n; i++ {
		row := []string{}
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	row := 0
	backtrack(board, row, &res)
	return res
}

func backtrack(board [][]string, row int, res *int) {
	if row == len(board) {
		(*res)++
		return
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = "Q"
			backtrack(board, row+1, res)
			board[row][col] = "."
		}
	}
}

func isValid(board [][]string, row int, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == "Q" {
			return false
		}
	}
	//check右上角
	m, n := row-1, col+1
	for m >= 0 && n < len(board) {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n++
	}
	//check左上角
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
