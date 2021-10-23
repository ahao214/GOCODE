package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 1609 · 链表的中间结点
 * @param head: the head node
 * @return: the middle node
 */
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
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length%2 == 0 {
		return slow.Next
	}
	return slow
}
