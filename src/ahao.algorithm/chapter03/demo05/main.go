package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//判断两棵二叉树是否相等
//参数：root1和root2分别为两棵二叉树的根结点
//返回值：true-如果两棵树相等则返回true，否则返回false
func IsEqual(root1 *gotype.BNode, root2 *gotype.BNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1.Data == root2.Data {
		return IsEqual(root1.LeftChild, root2.LeftChild) && IsEqual(root1.RightChild, root2.RightChild)
	}
	return false
}

func CreateTree() *gotype.BNode {
	root := &gotype.BNode{}
	node1 := &gotype.BNode{}
	node2 := &gotype.BNode{}
	node3 := &gotype.BNode{}
	node4 := &gotype.BNode{}

	root.Data = 6
	node1.Data = 3
	node2.Data = -7
	node3.Data = -1
	node4.Data = 9
	root.LeftChild = node1
	root.RightChild = node2
	node1.LeftChild = node3
	node1.RightChild = node4
	return root

}

func CreateTree2() *gotype.BNode {
	root := &gotype.BNode{}
	node1 := &gotype.BNode{}
	node2 := &gotype.BNode{}
	node3 := &gotype.BNode{}
	node4 := &gotype.BNode{}

	root.Data = 9
	node1.Data = 3
	node2.Data = -7
	node3.Data = -1
	node4.Data = 9
	root.LeftChild = node1
	root.RightChild = node2
	node1.LeftChild = node3
	node1.RightChild = node4
	return root

}

//判断两棵二叉树是否相等
func main() {
	root1 := CreateTree()
	root2 := CreateTree2()
	if IsEqual(root1, root2) {
		fmt.Println("这两棵树相等")
	} else {
		fmt.Println("这两棵树不相等")
	}
}
