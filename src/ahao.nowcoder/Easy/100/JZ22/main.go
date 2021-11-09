package main

//JZ22 链表中倒数最后k个结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	fast := pHead
	slow := pHead
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
