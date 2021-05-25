package main

import (
	"fmt"
)

//求解迷宫问题
var N = 4

//打印从起点到终点的路线
func printSolution(sol [][]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(sol[i][j], " ")
		}
		fmt.Println()
	}
}

//判断x和y是否是一个合理的单元
func isSafe(maze [][]int, x, y int) bool {
	return (x >= 0 && x < N && y >= 0 && y < N && maze[x][y] == 1)
}

//使用回溯的方法找到一条从左上角到右下角的路径
//maze表示迷宫，x/y表示起点，sol存储结果
func getPath(maze [][]int, x, y int, sol [][]int) bool {
	//到达目的地
	if x == N-1 && y == N-1 {
		sol[x][y] = 1
		return true
	}
	//判断maze[x][y]是否是一个可通行的单元
	if isSafe(maze, x, y) {
		//标记当前单元为1
		sol[x][y] = 1
		//向右走一步
		if getPath(maze, x+1, y, sol) {
			return true
		}

		//向下走一步
		if getPath(maze, x, y+1, sol) {
			return true
		}

		//标记当前单元为0用来表示这条路不可行，然后回溯
		sol[x][y] = 0
	}
	return false
}

func main() {
	maze := [][]int{{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1}}

	sol := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if getPath(maze, 0, 0, sol) {
		printSolution(sol)
	} else {
		fmt.Println("不存在可达的路径")
	}
}
