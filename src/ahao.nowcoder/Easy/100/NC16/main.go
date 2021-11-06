package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * NC16 判断二叉树是否对称
 * @param root TreeNode类
 * @return bool布尔型
 */
func isSymmetric(root *TreeNode) bool {
	return isEqual(root, root)
}

func isEqual(rootLeft *TreeNode, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return true
	}
	if rootLeft == nil || rootRight == nil || rootLeft.Val != rootRight.Val {
		return false
	}
	isLeftEqual := isEqual(rootLeft.Left, rootRight.Right)
	isRightEqual := isEqual(rootLeft.Right, rootRight.Left)
	return isLeftEqual && isRightEqual
}
