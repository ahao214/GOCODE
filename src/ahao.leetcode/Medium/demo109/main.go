package main

//109. 有序链表转换二叉搜索树
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		root := &TreeNode{}
		root.Val = head.Val
		return root
	}
	cur := head
	slow, fast, tmp := cur, cur, cur
	for fast != nil && fast.Next != nil {
		tmp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	nex := slow.Next
	slow.Next = nil
	tmp.Next = nil
	root := &TreeNode{}
	root.Val = slow.Val
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(nex)
	return root
}
