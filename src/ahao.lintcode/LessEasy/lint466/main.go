package main

//466 · 链表节点计数

type ListNode struct {
	Val  int
	Next *ListNode
}

func countNodes(head *ListNode) int {
	res := 0
	if head == nil {
		return res
	}
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
