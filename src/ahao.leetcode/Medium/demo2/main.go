package main

/*
2.两数相加
2个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点
输入：l1 = [2,4,3], l2 = [5,6,4]
解释：342 + 465 = 807
输出：[7,0,8]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	one, two, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			one = 0
		} else {
			one = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			two = 0
		} else {
			two = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (one + two + carry) % 10}
		current = current.Next
		carry = (one + two + carry) / 10
	}
	return head.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		sum := a + b + carry
		carry = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		cur.Next = &ListNode{carry, nil}
	}
	return pre.Next
}
