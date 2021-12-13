package main

/*
807. 保持城市天际线
在二维数组grid中(二维数组是n*n)，grid[i][j]代表位于某处的建筑物的高度。
我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。
高度 0 也被认为是建筑物。
最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。
城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。
建筑物高度可以增加的最大总和是多少？
*/

//计算出每一行每一列的最大值
func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)
	for i, row := range grid {
		for j, h := range row {
			rowMax[i] = max(rowMax[i], h)
			colMax[j] = max(colMax[j], h)
		}
	}

	for i, row := range grid {
		for j, h := range row {
			res += min(rowMax[i], colMax[j]) - h
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
