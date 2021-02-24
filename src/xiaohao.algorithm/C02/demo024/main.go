package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//814. 二叉树剪枝
func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

//递归求解
func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = deal(node.Left)
	node.Right = deal(node.Right)
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
	return node
}
