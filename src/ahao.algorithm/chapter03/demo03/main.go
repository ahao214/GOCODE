package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//方法功能：用层序遍历的方法打印出二叉树结点的内容
//输入参数：root-二叉树根结点
func PrintTreeLayer(root *gotype.BNode) {
	if root == nil {
		return
	}
	var p *gotype.BNode
	queue := gotype.NewSliceQueue()
	//树根结点进行队列
	queue.EnQueue(root)
	for queue.Size() > 0 {
		p = queue.DeQueue().(*gotype.BNode)
		//访问当前结点
		fmt.Print(p.Data, " ")
		//如果这个结点的左孩子不为空则入队列
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}
		//如果这个结点的有孩子不为空则入队列
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}

func printAtLevel(root *gotype.BNode, level int) int {
	if root == nil || level < 0 {
		return 0
	} else if level == 0 {
		fmt.Println(root.Data)
		return 1
	} else {
		//把打印根结点level层的结点转换为求解根结点的孩子结点的level-1
		//层的结点
		return printAtLevel(root.LeftChild, level-1) + printAtLevel(root.RightChild, level-1)
	}

}

//从顶部开始逐层打印二叉树结点数据
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("数组：", data)
	root := gotype.ArrayToTree(data, 0, len(data)-1)
	fmt.Print("树的层序遍历结果为：")
	PrintTreeLayer(root)
	fmt.Print()
}
