package NC70

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * NC70 单链表的排序
 * @param head ListNode类 the head node
 * @return ListNode类
 */
func sortInList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return head
}
