package main

//1036. 逃离大迷宫
/*
在一个 106 x 106 的网格中，每个网格上方格的坐标为 (x, y) 。

现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。

每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。

只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。
*/

//方法一：有限步数的广度优先搜索
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(block [][]int, source, target []int) bool {
	const (
		blocked = -1 // 在包围圈中
		valid   = 0  // 不在包围圈中
		found   = 1  // 无论在不在包围圈中，但在 n(n-1)/2 步搜索的过程中经过了 target

		boundary int = 1e6
	)

	n := len(block)
	if n < 2 {
		return true
	}

	blockSet := map[pair]bool{}
	for _, b := range block {
		blockSet[pair{b[0], b[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		countdown := n * (n - 1) / 2

		q := []pair{{sx, sy}}
		vis := map[pair]bool{{sx, sy}: true}
		for len(q) > 0 && countdown > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				np := pair{x, y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
					if x == fx && y == fy {
						return found
					}
					countdown--
					vis[np] = true
					q = append(q, np)
				}
			}
		}

		if countdown > 0 {
			return blocked
		}
		return valid
	}

	res := check(source, target)
	return res == found || res == valid && check(target, source) != blocked
}

//方法二：离散化 + 广度优先搜索
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 离散化 a，返回的哈希表中的键值对分别为 a 中的原始值及其离散化后的对应值
func discrete(a []int) (map[int]int, int) {
	sort.Ints(a)

	id := 0
	if a[0] > 0 {
		id = 1
	}
	mapping := map[int]int{a[0]: id}
	pre := a[0]
	for _, v := range a[1:] {
		if v != pre {
			if v == pre+1 {
				id++
			} else {
				id += 2
			}
			mapping[v] = id
			pre = v
		}
	}

	const boundary int = 1e6
	if a[len(a)-1] != boundary-1 {
		id++
	}

	return mapping, id
}

func isEscapePossible2(block [][]int, source, target []int) bool {
	n := len(block)
	if n < 2 {
		return true
	}
	rows := []int{source[0], target[0]}
	cols := []int{source[1], target[1]}
	for _, b := range block {
		rows = append(rows, b[0])
		cols = append(cols, b[1])
	}

	// 离散化行列坐标
	rMapping, rBound := discrete(rows)
	cMapping, cBound := discrete(cols)

	grid := make([][]bool, rBound+1)
	for i := range grid {
		grid[i] = make([]bool, cBound+1)
	}
	for _, b := range block {
		grid[rMapping[b[0]]][cMapping[b[1]]] = true
	}

	sx, sy := rMapping[source[0]], cMapping[source[1]]
	tx, ty := rMapping[target[0]], cMapping[target[1]]
	grid[sx][sy] = true
	q := []pair{{sx, sy}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x <= rBound && 0 <= y && y <= cBound && !grid[x][y] {
				if x == tx && y == ty {
					return true
				}
				grid[x][y] = true
				q = append(q, pair{x, y})
			}
		}
	}
	return false
}
