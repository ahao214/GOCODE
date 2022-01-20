package main

/*
328. 奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{Val: 0, Next: nil}
	odd := oddHead
	evenHead := &ListNode{Val: 0, Next: nil}
	even := evenHead

	count := 1
	for head != nil {
		if count%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		count++
	}
	even.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
}

func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//define odd even evenHead
	odd, even, evenHead := head, head.Next, head.Next
	for even != nil && even.Next != nil {
		oddNext := even.Next
		odd.Next = oddNext
		odd = odd.Next
		evenNext := odd.Next
		even.Next = evenNext
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
