package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//把一个有序整数数组放到二叉树中
func arrayToTree(arr []int, start int, end int) *gotype.BNode {
	var root *gotype.BNode
	if end >= start {
		root = gotype.NewBNode()
		mid := (start + end + 1) / 2
		//树的根结点为数组中间的元素
		root.Data = arr[mid]
		//递归的用左半部分数组构造root的左子树
		root.LeftChild = arrayToTree(arr, start, mid-1)
		//递归的用有半部分数组构造root的右子树
		root.RightChild = arrayToTree(arr, mid+1, end)

	}
	return root
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("数组：", data)
	root := arrayToTree(data, 0, len(data)-1)
	fmt.Print("转换成树的中序遍历为：")
	gotype.PrintTreeMidOrder(root)
	fmt.Print()
}
