package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//创建链表
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

//判断两个链表(无环)是否交叉
func main() {
	fmt.Println("判断相交链表")

}
