package main

/*
876. 链表的中间结点
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length%2 == 0 {
		return slow.Next
	}
	return slow
}

func middleNode1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil && fast.Next.Next == nil {
		slow = slow.Next
	}
	return slow
}
