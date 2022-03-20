package demo21

/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	cur := pre
	for list1 != nil && list2 != nil {
		tmp := list1.Val
		if list1.Val > list2.Val {
			tmp = list2.Val
		}
		cur.Next = &ListNode{tmp, nil}
		cur = cur.Next
		if list1.Val > list2.Val {
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return pre.Next
}
