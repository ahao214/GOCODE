package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * NC96 判断一个链表是否为回文结构
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	if head == nil {
		return true
	}
	lst := []int{}
	for head != nil {
		lst = append(lst, head.Val)
		head = head.Next
	}
	for i, j := 0, len(lst)-1; i < j; {
		if lst[i] != lst[j] {
			return false
		}
		i++
		j--
	}
	return true
}
