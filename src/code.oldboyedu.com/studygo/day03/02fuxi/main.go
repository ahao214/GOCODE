package main

import (
	"fmt"

	"github.com/isdamir/gotype"

)


//带头节点的逆序
func Reverse(Node *LNode){
	if node==nil || node.Next==nil {
		return
	}

	var pre *LNode
	var cur *LNode
	next:=node.Next
	for next!=nil{
		cur = next.Next
		next.Next =pre
		pre=next
		next=cur
	}
	node.Next=pre
}
func main(){
	head:=&LNode{}
	
	fmt.Println("就地逆序")
	CreateNode(head, 8)
	PrintNode("逆序前："，head)
	Reverse(head)
	PrintNode("逆序后：",head)


}