package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func Reverse(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head.Next       //当前遍历结点
	pre := head            //当前结点的前驱结点
	var next *gotype.LNode //当前结点后继结点的后继结点
	for cur != nil && cur.Next != nil {
		next = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
}

//把链表相邻元素翻转
func main() {
	fmt.Println("相邻元素翻转")
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	gotype.PrintNode("顺序输出：", head)
	Reverse(head)
	gotype.PrintNode("逆序输出：", head)

}
