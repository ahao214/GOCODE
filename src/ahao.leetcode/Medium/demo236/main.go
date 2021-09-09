package main

//236. 二叉树的最近公共祖先
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestTreeNode(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestTreeNode(root.Left, p, q)
	right := LowestTreeNode(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
