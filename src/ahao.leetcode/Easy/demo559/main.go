package main

//559. N 叉树的最大深度
type Node struct {
	Val      int
	Children []*Node
}

//深度优先搜索
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth
}
