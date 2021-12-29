package main

//129. 求根节点到叶节点数字之和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	return helper(root, 0)
}

func helper(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	num = root.Val + num*10
	if root.Left == nil && root.Right == nil {
		return num
	}
	sum := 0
	sum += helper(root.Left, num)
	sum += helper(root.Right, num)
	return sum
}
