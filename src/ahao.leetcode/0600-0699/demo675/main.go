package main

import "sort"

/*
675. 为高尔夫比赛砍树
你被请来给一个要举办高尔夫比赛的树林砍树。树林由一个 m x n 的矩阵表示， 在这个矩阵中：

0 表示障碍，无法触碰
1 表示地面，可以行走
比 1 大的数 表示有树的单元格，可以行走，数值表示树的高度
每一步，你都可以向上、下、左、右四个方向之一移动一个单位，如果你站的地方有一棵树，那么你可以决定是否要砍倒它。

你需要按照树的高度从低向高砍掉所有的树，每砍过一颗树，该单元格的值变为 1（即变为地面）。

你将从 (0, 0) 点开始工作，返回你砍完所有树需要走的最小步数。 如果你无法砍完所有的树，返回 -1 。

可以保证的是，没有两棵树的高度是相同的，并且你至少需要砍倒一棵树。



示例 1：


输入：forest = [[1,2,3],[0,0,4],[7,6,5]]
输出：6
解释：沿着上面的路径，你可以用 6 步，按从最矮到最高的顺序砍掉这些树。
示例 2：


输入：forest = [[1,2,3],[0,0,0],[7,6,5]]
输出：-1
解释：由于中间一行被障碍阻塞，无法访问最下面一行中的树。
示例 3：

输入：forest = [[2,3,4],[0,0,5],[8,7,6]]
输出：6
解释：可以按与示例 1 相同的路径来砍掉所有的树。
(0,0) 位置的树，可以直接砍去，不用算步数。


提示：

m == forest.length
n == forest[i].length
1 <= m, n <= 50
0 <= forest[i][j] <= 109
*/

//广度优先
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func cutOffTree(forest [][]int) (ans int) {
	type pair struct{ dis, x, y int }
	trees := []pair{}
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, pair{h, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })

	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}

	preX, preY := 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return
}
