package main

//NC102 在二叉树中找到两个节点的最近公共祖先
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	if root == nil {
		return -1
	}
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}
	left := lowestCommonAncestor(root.Left, o1, o2)
	right := lowestCommonAncestor(root.Right, o1, o2)
	if left != 0 && right != 0 {
		return root.Val
	}
	if left == 0 {
		return right
	}
	return left
}
