package main

//JZ52 两个链表的第一个公共结点
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	slow := pHead1
	fast := pHead2
	for slow != fast {
		if slow == nil {
			slow = pHead2
		} else {
			slow = slow.Next
		}
		if fast == nil {
			fast = pHead1
		} else {
			fast = fast.Next
		}
	}
	return fast
}
