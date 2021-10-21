package main

type Node struct {
	Val      int
	Children []*Node
}

//429. N 叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		length := len(q)
		level := []int{}
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			for _, v := range node.Children {
				q = append(q, v)
			}
		}
		if len(level) > 0 {
			res = append(res, level)
		}
	}
	return res
}
