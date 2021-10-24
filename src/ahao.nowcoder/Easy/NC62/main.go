package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * NC62 平衡二叉树
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return abs(height(pRoot.Left)-height(pRoot.Right)) <= 1 && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
