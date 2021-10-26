package main

//994. 腐烂的橘子
func orangesRotting(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	res := 0
	for {
		flag := true
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				// 如果橘子坏了，会感染旁边的好橘子
				if grid[i][j] == 2 {
					if i > 0 {
						if grid[i-1][j] == 1 {
							grid[i-1][j] = 3
							flag = false
						}
					}
					if j > 0 {
						if grid[i][j-1] == 1 {
							grid[i][j-1] = 3
							flag = false
						}
					}
					if i < row-1 {
						if grid[i+1][j] == 1 {
							grid[i+1][j] = 3
							flag = false
						}
					}
					if j < col-1 {
						if grid[i][j+1] == 1 {
							grid[i][j+1] = 3
							flag = false
						}
					}
				}
			}
		}
		// 如果本次没有任何橘子腐烂，那么表示没有好橘子或者有好橘子对怀橘子不可达
		if flag {
			break
		}

		res++
		// 把刚刚标记为 3 的橘子改为 2
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}
	}

	if hasFresh(grid) {
		return -1
	}

	return res
}

//判断是否有新鲜的橘子
func hasFresh(grid [][]int) bool {
	for _, row := range grid {
		for _, col := range row {
			if col == 1 {
				return true
			}
		}
	}
	return false
}
