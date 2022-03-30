package main

//563. 二叉树的坡度
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res = 0
	nodeTilt(root)
	return res
}

var res int

func nodeTilt(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sumLeft := nodeTilt(node.Left)
	sumRight := nodeTilt(node.Right)
	res += abs(sumLeft - sumRight)
	return sumLeft + sumRight + node.Val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
