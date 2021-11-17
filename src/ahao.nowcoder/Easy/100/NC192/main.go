package main

//NC192 二叉树的后序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	postOrder(root, &result)
	return result
}

func postOrder(root *TreeNode, result *[]int) {
	if root != nil {
		postOrder(root.Left, result)
		postOrder(root.Right, result)
		*result = append(*result, root.Val)
	}
}
