package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//给定单链表中某个结点，删除该结点
func RemoveNode(node *gotype.LNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Data = node.Next.Data
	tmp := node.Next
	node.Next = tmp.Next
	return true
}

//在只给定单链表中某个结点指针的情况下删除该结点
func main() {
	fmt.Println("删除指定结点")
	head := &gotype.LNode{}
	retNode := CreateNodeT(head, 5)
	fmt.Printf("删除节点%v之前链表是\n", retNode.Data)
	gotype.PrintNode("链表：", head)
	result := RemoveNode(retNode)
	if result {
		gotype.PrintNode("删除该结点后链表：", head)
	}
}

//创建单链表
func CreateNodeT(node *gotype.LNode, get int) (retNode *gotype.LNode) {
	cur := node
	for i := 1; i < 8; i++ {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i
		cur = cur.Next
		if i == get {
			retNode = cur
		}
	}
	return
}
