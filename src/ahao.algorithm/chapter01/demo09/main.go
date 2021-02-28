package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//合并两个升序排列的单链表
func Merge(head1 *gotype.LNode, head2 *gotype.LNode) *gotype.LNode {
	if head1 == nil || head1.Next == nil {
		return head1
	}
	if head2 == nil || head2.Next == nil {
		return head2
	}
	cur1 := head1.Next     //用来遍历head1
	cur2 := head2.Next     //用来遍历head2
	var head *gotype.LNode //合并后链表的头结点
	var cur *gotype.LNode  //合并后的链表在尾结点
	if cur1.Data.(int) > cur2.Data.(int) {
		head = head2
		cur = cur2
		cur2 = cur2.Next
	} else {
		head = head1
		cur = cur1
		cur1 = cur1.Next
	}
	//每次找链表剩余结点的最小值对应的结点连接到合并后链表的尾部
	for cur1 != nil && cur2 != nil {
		if cur1.Data.(int) < cur2.Data.(int) {
			cur.Next = cur1
			cur = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur = cur2
			cur2 = cur2.Next
		}
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return head
}

//合并两个有序链表
func main() {
	fmt.Println("合并有序链表")
	head1 := &gotype.LNode{}
	head2 := &gotype.LNode{}
	CreateNodeT(head1, 1)
	CreateNodeT(head2, 2)
	gotype.PrintNode("head1:", head1)
	gotype.PrintNode("head2:", head2)
	head := Merge(head1, head2)
	gotype.PrintNode("合并后的链表：", head)
}

func CreateNodeT(node *gotype.LNode, start int) {
	cur := node
	for i := start; i < 7; i += 2 {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}
