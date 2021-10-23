package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * NC33 合并两个排序的链表
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	if pHead1.Val < pHead2.Val {
		pHead1.Next = Merge(pHead1.Next, pHead2)
		return pHead1
	}
	pHead2.Next = Merge(pHead2.Next, pHead1)
	return pHead2
}
