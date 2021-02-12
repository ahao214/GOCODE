package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//顺序删除
func RemoveDup(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}

	outerCur := head.Next      //用于外出循环，指向链表第一个结点
	var innerCur *gotype.LNode //用于内层循环遍历outerCur后面的结点
	var innerPre *gotype.LNode //innerCur的前驱结点
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

//从无序链表中移除重复项
func main() {
	head := &gotype.LNode{}
	fmt.Println("删除重复结点")
	gotype.CreateNode(head, 8)
	fmt.Println("---顺序删除---")
	gotype.PrintNode("删除重复结点前：", head)
	RemoveDup(head)
	gotype.PrintNode("删除重复结点后：", head)
}
