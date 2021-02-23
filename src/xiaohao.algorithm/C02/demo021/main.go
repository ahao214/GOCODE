package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//到这里意味着已经查找到目标
	if root.Right == nil {
		//右子树为空
		return root.Left
	}
	if root.Left == nil {
		//左子树为空
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil {
		//查找后续
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right)
	return root
}

func deleteMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteMinNode(root.Left)
	return root
}

//第450题：删除二叉搜索树中的节点
func main() {

}
