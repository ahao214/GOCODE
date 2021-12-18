package main

//419.甲板上的战舰

//遍历扫描
func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	ans := 0
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				row[j] = '.'
				for k := j + 1; k < n && row[k] == 'X'; k++ {
					row[k] = '.'
				}
				for k := i + 1; k < m && board[k][j] == 'X'; k++ {
					board[k][j] = '.'
				}
				ans++
			}
		}
	}
	return ans
}

//枚举起点
func countBattleships2(board [][]byte) int {
	ans := 0
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' && !(i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X') {
				ans++
			}
		}
	}
	return ans
}
