package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 反转链表
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	newHead := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return newHead
}
