package main

/*
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	len := 0
	cur := newHead
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	if (k % len) == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cur := head
	count := 0
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	count++
	//tail next connect to the head
	cur.Next = head
	k = k % count
	n := count - k
	for n > 1 {
		head = head.Next
		n--
	}
	newHead := head.Next
	head.Next = nil
	return newHead
}
