package main

//66 · 二叉树的前序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: A Tree
 * @return: Preorder in ArrayList which contains node values.
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	preOrder(root, &res)
	return res
}

func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
