package main

//1219. 黄金矿工
func getMaximumGold(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	dir := make([][]int, 4)
	dir[0] = []int{1, 0}
	dir[1] = []int{0, 1}
	dir[2] = []int{-1, 0}
	dir[3] = []int{0, -1}
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				backtrack(grid, i, j, dir, 0, &maxGold)
			}
		}
	}
	return maxGold
}

func backtrack(grid [][]int, n int, m int, dir [][]int, res int, maxGold *int) int {
	res += grid[n][m]
	if res > *maxGold {
		*maxGold = res
	}
	tmp := grid[n][m]
	grid[n][m] = 0
	for index := 0; index < len(dir); index++ {
		i := n + dir[index][0]
		j := m + dir[index][1]
		//超过单元格或者黄金数量为0
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
			continue
		}
		backtrack(grid, i, j, dir, res, maxGold)
	}
	grid[n][m] = tmp
	return res
}
