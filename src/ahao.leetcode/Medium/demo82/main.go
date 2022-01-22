package main

/*
82. 删除排序链表中的重复元素 II
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，
只保留原始链表中 没有重复出现 的数字。

返回同样按升序排列的结果链表。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	// 上次遍历有删除操作的标志位
	lastIsDel := false
	// 虚拟空结点
	head = nilNode
	// 前后指针用于判断
	pre, back := head.Next, head.Next.Next
	// 每次只删除前面的一个重复的元素，留一个用于下次遍历判重
	// pre, back 指针的更新位置和值比较重要和巧妙
	for head.Next != nil && head.Next.Next != nil {
		if pre.Val != back.Val && lastIsDel {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
			continue
		}

		if pre.Val == back.Val {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = true
		} else {
			head = head.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
		}
	}
	// 处理 [1,1] 这种删除还剩一个的情况
	if lastIsDel && head.Next != nil {
		head.Next = nil
	}
	return nilNode.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
