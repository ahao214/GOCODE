package main

import (
	"fmt"
	"github.com/isdamir/gotype" //引入定义的数据结构
)

//就地逆序
//带头结点的逆序
func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *gotype.LNode //定义前驱结点
	var cur *gotype.LNode //定义当前结点
	next := node.Next     //后续结点
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

//递归法
func RecursiveReverseChild(node *gotype.LNode) *gotype.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *gotype.LNode) {
	firstNode := node.Next
	//递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

//插入法
func InsertReverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *gotype.LNode  //前序结点
	var next *gotype.LNode //后续结点
	cur = node.Next.Next
	//设置链表第一个结点为尾结点
	node.Next.Next = nil
	//把遍历到的结点插入到头结点的前面
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}

}

//实现链表的逆序
func main() {
	head := &gotype.LNode{}
	fmt.Println("就地逆序")
	gotype.CreateNode(head, 8)
	gotype.PrintNode("逆序前：", head)
	Reverse(head)
	gotype.PrintNode("逆序后：", head)

}
