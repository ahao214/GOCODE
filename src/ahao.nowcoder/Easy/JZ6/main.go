package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * JZ6 从尾到头打印链表
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
