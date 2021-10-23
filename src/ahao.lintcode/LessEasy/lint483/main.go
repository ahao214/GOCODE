package main

//483 · 链表转数组
type ListNode struct {
	Val  int
	Next *ListNode
}

func toArrayList(head *ListNode) []int {
	res := []int{}
	if head == nil {
		return res
	}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
