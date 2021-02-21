package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//找出链表Head的中间结点，把链表从中间断成两个子链表
func findMiddleNode(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head //遍历链表的每次向前走两步
	slow := head //遍历链表的每次向前走一步

	slowPre := head
	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil //分割链表为两部分
	return slow
}

//对不带头结点的单链表翻转
func reverse(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *gotype.LNode  //前驱结点
	var next *gotype.LNode //后驱结点
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func Reorder(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}

	cur1 := head.Next //前半部分链表的第一个结点
	mid := findMiddleNode(head.Next)
	cur2 := reverse(mid) //后半部分链表逆序后的第一个结点
	var tmp *gotype.LNode
	//合并链表
	for cur1.Next != nil {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2
}

//对链表进行重新排序
func main() {
	fmt.Println("链表排序")
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	gotype.PrintNode("排序前：", head)
	Reorder(head)
	gotype.PrintNode("排序后：", head)

}
