package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func Add(h1 *gotype.LNode, h2 *gotype.LNode) *gotype.LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	c := 0   //记录进位
	sum := 0 //记录两个结点相加的值
	p1 := h1.Next
	p2 := h2.Next
	resultHead := &gotype.LNode{}
	p := resultHead
	for p1 != nil && p2 != nil {
		p.Next = &gotype.LNode{} //指向新创建的存储相加和的结点
		sum = p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum % 10 //两结点相加和
		c = sum / 10           //进位
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	//链表h2比h1长，接下来只需要考虑h2剩余结点的值
	if p1 == nil {
		for p2 != nil {
			p.Next = &gotype.LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}

	//链表h1比h2长，接下来只需要考虑h1剩余结点的值
	if p2 == nil {
		for p1 != nil {
			p.Next = &gotype.LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}

	if c == 1 {
		p.Next = &gotype.LNode{}
		p.Next.Data = 1
	}
	return resultHead
}

//创建链表
func CreateNodeT() (l1 *gotype.LNode, l2 *gotype.LNode) {
	l1 = &gotype.LNode{}
	l2 = &gotype.LNode{}
	cur := l1
	for i := 1; i < 7; i++ {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}

	cur = l2
	for i := 9; i > 4; i-- {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
	return
}

//计算两个单链表所代表的数之和
func main() {
	fmt.Println("链表相加")
	l1, l2 := CreateNodeT()
	gotype.PrintNode("head1:", l1)
	gotype.PrintNode("head2:", l2)

	addResult := Add(l1, l2)
	gotype.PrintNode("相加后：", addResult)

}
