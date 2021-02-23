package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归方法
func countNodesDG(root *TreeNode) int {
	if root != nil {
		return 1 + countNodesDG(root.Left) + countNodesDG(root.Right)
	}
	return 1 + countNodesDG(root.Left) + countNodesDG(root.Right)
}

//第222题：完全二叉树的节点个数
//完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大
//值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个
//节点
func main() {

}
