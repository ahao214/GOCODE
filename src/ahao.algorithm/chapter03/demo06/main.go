package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

var pHead *gotype.BNode
var pEnd *gotype.BNode

func InOrderBSTree(root *gotype.BNode) {
	if root == nil {
		return
	}
	//转换root的左子树
	InOrderBSTree(root.LeftChild)
	//使当前结点的左孩子指向双向链表中最后一个结点
	root.LeftChild = pEnd
	//双向列表为空，当前遍历的结点为双向链表的头结点
	if pEnd == nil {
		pHead = root
	} else {
		//使双向链表中最后一个结点的右孩子指向当前结点
		pEnd.RightChild = root
	}
	pEnd = root //将当前结点设为双向链表中最后一个结点
	//转换root的右子树
	InOrderBSTree(root.RightChild)

}

//把二叉树转换为双向链表
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("数组：", data)
	root := gotype.ArrayToTree(data, 0, len(data)-1)
	InOrderBSTree(root)
	fmt.Print("转换后双向链表正向遍历：")
	for cur := pHead; cur != nil; cur = cur.RightChild {
		fmt.Print(cur.Data, " ")
	}
	fmt.Print()
	fmt.Print("转换后双向链表逆向遍历：")
	for cur := pEnd; cur != nil; cur = cur.LeftChild {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}
