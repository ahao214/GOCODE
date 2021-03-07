package main

import (
	"fmt"
	"github.com/isdamir/gotype"
	"math"
)

var maxSum = math.MinInt64

//求最大子树
func FindMaxSubTree(root *gotype.BNode, maxRoot *gotype.BNode) int {
	if root == nil {
		return 0
	}

	//求root左子树所有结点的和
	lmax := FindMaxSubTree(root.LeftChild, maxRoot)
	//求root右子树所有结点的和
	rmax := FindMaxSubTree(root.RightChild, maxRoot)
	sum := lmax + rmax + root.Data.(int)
	//以root为根的子树的和大于前面求出的最大值
	if sum > maxSum {
		maxSum = sum
		maxRoot.Data = root.Data
	}
	//返回以root为根结点的子树的所有结点的和
	return sum
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

//求一颗二叉树的最大子树和
func main() {
	//构造二叉树
	root := CreateTree()
	//最大子树的根结点
	maxRoot := &gotype.BNode{}
	FindMaxSubTree(root, maxRoot)
	fmt.Println("最大子树和为：", maxSum)
	fmt.Println("对应子树的根结点为：", maxRoot.Data)

}
