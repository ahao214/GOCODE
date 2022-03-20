package main

//145. 二叉树的后序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, res)
	postOrder(root.Right, res)
	*res = append(*res, root.Val)
}
