package main

//JZ18 删除链表的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := &ListNode{0, head}
	pre.Next = head
	cur := head
	next := &ListNode{}
	for cur != nil {
		cur = cur.Next
		pre = pre.Next
		if cur.Val == val {
			next = cur.Next
			break
		}
	}

	cur.Next = nil
	pre.Next = next
	return head
}
