package main

//450.删除二叉搜索树中的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		left := root.Left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = root.Right
		root = root.Left
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
