package main

/*
143. 重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//cut in half
	second := slow.Next
	slow.Next = nil
	//reverse second half
	var pre *ListNode
	for second != nil {
		nxt := second.Next
		second.Next = pre
		pre = second
		second = nxt
	}
	//insert node 2 to node 1
	for pre != nil {
		nxt := head.Next
		preNext := pre.Next

		head.Next = pre
		pre.Next = nxt

		head = nxt
		pre = preNext
	}
}
