package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//对不带头结点的单链表翻转
func Reverse(head *gotype.LNode) *gotype.LNode {
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

func ReverseK(head *gotype.LNode, k int) {
	if head == nil || head.Next == nil {
		return
	}
	pre := head
	begin := head.Next
	var end *gotype.LNode
	var pNext *gotype.LNode
	for begin != nil {
		end = begin
		//找到从begin开始第k个结点
		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				return
			}
		}
		pNext = end.Next
		end.Next = nil
		pre.Next = Reverse(begin)
		begin.Next = pNext
		pre = begin
		begin = pNext
	}
}

//把链表以K个结点为一组进行翻转
func main() {
	fmt.Println("k元素翻转")
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	gotype.PrintNode("顺序输出：", head)
	ReverseK(head, 3)
	gotype.PrintNode("逆序输出：", head)
}
