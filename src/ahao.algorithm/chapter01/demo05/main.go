package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//快慢指针查找
func FindLastK(head *gotype.LNode, k int) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next
	i := 0
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return nil
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow

}

//找出单链表中的倒数第K个元素
func main() {
	fmt.Println("寻找倒数K")
	head := &gotype.LNode{}
	gotype.CreateNode(head, 9)
	gotype.PrintNode("链表：", head)
	fmt.Println("链表倒数第3个元素为：", FindLastK(head, 3).Data)
}
