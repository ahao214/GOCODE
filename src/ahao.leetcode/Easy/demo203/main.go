package main

//203.移除链表元素
//给你一个链给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElement(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElement(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
