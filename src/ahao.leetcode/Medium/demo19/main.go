package demo19

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
19. 删除链表的倒数第N个结点
说明：给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
