package main

//572. 另一棵树的子树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(b.Left, b.Right)
	}

	return false
}
