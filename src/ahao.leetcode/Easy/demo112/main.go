package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//112. 路径总和
func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return pathSum(root.Left, sum-root.Val) || pathSum(root.Right, sum-root.Val)

}
