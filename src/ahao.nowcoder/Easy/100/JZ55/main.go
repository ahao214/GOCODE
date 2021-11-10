package main

//JZ55 二叉树的深度
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return max(TreeDepth(pRoot.Left), TreeDepth(pRoot.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
