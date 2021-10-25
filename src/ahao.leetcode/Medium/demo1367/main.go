package main

//1367. 二叉树中的列表
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return helper(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func helper(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil && head != nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return helper(head.Next, root.Left) || helper(head.Next, root.Right)
}
