package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * NC45 实现二叉树先序，中序和后序遍历
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	res := [][]int{}
	pre := []int{}
	in := []int{}
	post := []int{}
	preOrder(root, &pre)
	inOrder(root, &in)
	postOrder(root, &post)
	res = append(res, pre, in, post)
	return res
}

//先序
func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

//中序
func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

//后序
func postOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, res)
	postOrder(root.Right, res)
	*res = append(*res, root.Val)
}
