package main

/*
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//find middle and reverse left part
	slow, fast := head, head
	var pre *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		nxt := slow.Next
		slow.Next = pre
		pre = slow
		slow = nxt
	}

	var left *ListNode
	var right *ListNode
	//define left and right pointers
	if fast.Next == nil {
		left = pre
		right = slow.Next
	} else {
		right = slow.Next
		slow.Next = pre
		left = slow
	}
	//compare
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
