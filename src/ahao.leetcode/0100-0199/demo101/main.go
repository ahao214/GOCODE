package main

/*
101. 对称二叉树
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isEqual(root.Left, root.Right)
}
func isEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	isLeftEqual := isEqual(left.Left, right.Right)
	isRightEqual := isEqual(left.Right, right.Left)
	return isLeftEqual && isRightEqual
}
