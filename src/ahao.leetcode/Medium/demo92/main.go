package main

/*
92. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	length := right - left
	var tail *ListNode
	if left > 1 {
		tail = &ListNode{-1, head}
	}
	cur := head
	for left > 1 {
		tail = tail.Next
		cur = cur.Next
		left--
	}
	newCur := cur
	//reverse nodes between left to right
	var pre *ListNode
	nxt := cur.Next

	for length >= 0 {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
		length--
	}
	//connect node1's tail to node2's head
	if tail != nil {
		tail.Next = pre
		newCur.Next = nxt
	} else {
		head = pre
		newCur.Next = nxt
	}
	return head
}
