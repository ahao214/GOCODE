package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * JZ28 对称的二叉树
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return isEqual(pRoot.Left, pRoot.Right)
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
