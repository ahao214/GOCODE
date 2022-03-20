package main

//100. 相同的树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return helper(p, q, true)
}

func helper(p *TreeNode, q *TreeNode, res bool) bool {
	if p == nil && q == nil {
		return res
	} else if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}
	return helper(p.Left, q.Left, res) && helper(p.Right, q.Right, res)
}
