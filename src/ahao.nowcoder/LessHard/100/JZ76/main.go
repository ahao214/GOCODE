package main

//JZ76 删除链表中重复的结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	dummy := &ListNode{0, pHead}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
