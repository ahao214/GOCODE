package main

//404. 左叶子之和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}
	helper(root.Left, res)
	helper(root.Right, res)
}
