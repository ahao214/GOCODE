package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//NC3 链表中环的入口结点
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	slow, fast := pHead, pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	third := pHead
	for third != slow {
		third = third.Next
		slow = slow.Next
	}
	return slow
}
