package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	inOrder(root, &sum)
	return root
}

func inOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	inOrder(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	inOrder(root.Left, sum)
}
