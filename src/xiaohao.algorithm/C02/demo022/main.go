package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	leftH := maxDepth(root.Left) + 1
	rightH := maxDepth(root.Right) + 1
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//第110题：平衡二叉树
//一棵高度平衡二叉树定义为：一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1
func main() {

}
