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

//广度优先搜索
func maxDepth2(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
		ans++
	}
	return ans
}
