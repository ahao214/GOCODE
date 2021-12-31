package main

//114. 二叉树展开为链表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			left := root.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
}
