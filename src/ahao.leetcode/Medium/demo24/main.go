package main

/*
24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		nxt := cur.Next
		nextNxt := nxt.Next

		pre.Next = nxt
		nxt.Next = cur
		cur.Next = nextNxt
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}
