package main

//1219. 黄金矿工
/*
你要开发一座金矿，地质勘测学家已经探明了这座金矿中的资源分布，并用大小为 m * n 的网格 grid 进行了标注。每个单元格中的整数就表示这一单元格中的黄金数量；如果该单元格是空的，那么就是 0。

为了使收益最大化，矿工需要按以下规则来开采黄金：

每当矿工进入一个单元，就会收集该单元格中的所有黄金。
矿工每次可以从当前位置向上下左右四个方向走。
每个单元格只能被开采（进入）一次。
不得开采（进入）黄金数目为 0 的单元格。
矿工可以从网格中 任意一个 有黄金的单元格出发或者是停止。
*/
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
