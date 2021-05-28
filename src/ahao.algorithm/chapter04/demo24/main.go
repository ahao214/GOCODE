package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//对有大量重复的数字的数组排序

//AVL树
type Sort struct{}

//中序遍历AVL树，把遍历结果放入到数组中
func (p *Sort) inorder(arr []int, root *AVLNode, index int) int {
	if root != nil {
		//中序遍历左子树
		index = p.inorder(arr, root.Left, index)
		for i := 0; i < root.Count; i++ {
			arr[index] = root.Data
			index++
		}
		//中序遍历右子树
		index = p.inorder(arr, root.Right, index)
	}
	return index
}

//得到树的高度
func (p *Sort) getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	} else {
		return node.Height
	}
}

//把以y为根的子树向右旋转
func (p *Sort) rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right
	//旋转
	x.Right = y
	y.Left = T2
	y.Height = Max(p.getHeight(y.Left), p.getHeight(y.Right))
	x.Height = Max(p.getHeight(x.Left), p.getHeight(x.Right))
	//返回新的根结点
	return x
}

//把以x为根的子树向左旋转
func (p *Sort) leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left
	//旋转
	y.Left = x
	x.Right = T2
	x.Height = Max(p.getHeight(x.Left), p.getHeight(x.Right))
	y.Height = Max(p.getHeight(y.Left), p.getHeight(y.Right))
	//返回新的根结点
	return y
}

//获取树的平衡因子
func (p *Sort) getBalance(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return p.getHeight(n.Left) - p.getHeight(n.Right)
}

//如果data在AVL树中不存在，则把data插入到ALV树中
//否则把这个结点对应的count加1即可
func (p *Sort) insert(root *AVLNode, data int) *AVLNode {
	if root == nil {
		return NewAVLNode(data)
	}
	//data在树中存在，把对应的结点的count加1
	if data == root.Data {
		root.Count += 1
		return root
	}
	if data < root.Data {
		//在左子树中继续查找data是否存在
		root.Left = p.insert(root.Left, data)
	} else {
		//在右子树中继续查找data是否存在
		root.Right = p.insert(root.Right, data)
	}
	//入新的结点后更新root结点的高度
	root.Height = Max(p.getHeight(root.Left), p.getHeight(root.Right)) + 1
	//获取树的平衡因子
	balance := p.getBalance(root)
	//如果树不平衡，根据数据结构中学过的四种情况进行调整
	if balance > 1 && data < root.Left.Data {
		//LL型
		return p.rightRotate(root)
	} else if balance < -1 && data > root.Right.Data {
		//RR型
		return p.leftRotate(root)
	} else if balance > 1 && data > root.Left.Data {
		//LR型
		root.Left = p.leftRotate(root.Left)
		return p.rightRotate(root)
	} else if balance < -1 && data < root.Right.Data {
		//RL型
		root.Right = p.rightRotate(root.Right)
		return p.leftRotate(root)
	}
	//返回数的根结点
	return root
}

//使用AVL树实现排序
func (p *Sort) Sort(arr []int) {
	//根结点
	var root *AVLNode
	for _, v := range arr {
		root = p.insert(root, v)
	}
	p.inorder(arr, root, 0)
}

func main() {
	arr := []int{15, 12, 15, 2, 2, 12, 2, 3, 12, 100, 3, 3}
	s := &Sort{}
	s.Sort(arr)
	fmt.Println("方法一：")
	for _, v := range arr {
		fmt.Print(v, " ")
	}

}
