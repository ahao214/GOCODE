package main

//110. 平衡二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(helper(root.Left, 0)-helper(root.Right, 0)) < 2 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func helper(root *TreeNode, height int) int {
	if root == nil {
		return height
	}
	return max(helper(root.Left, height+1), helper(root.Right, height+1))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
