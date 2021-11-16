package main

//NC193 二叉树的前序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	preOrder(root, &result)
	return result
}

func preOrder(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preOrder(root.Left, result)
		preOrder(root.Right, result)
	}
}
