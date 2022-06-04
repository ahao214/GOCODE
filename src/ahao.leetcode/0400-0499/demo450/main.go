package main

/*
450. 删除二叉搜索树中的节点
*/

//方法一：递归
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}


//方法二：迭代
func deleteNode2(root *TreeNode, key int) *TreeNode {
	var cur, curParent *TreeNode = root, nil
	for cur != nil && cur.Val != key {
		curParent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		cur = nil
	} else if cur.Right == nil {
		cur = cur.Left
	} else if cur.Left == nil {
		cur = cur.Right
	} else {
		successor, successorParent := cur.Right, cur
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		if successorParent.Val == cur.Val {
			successorParent.Right = successor.Right
		} else {
			successorParent.Left = successor.Right
		}
		successor.Right = cur.Right
		successor.Left = cur.Left
		cur = successor
	}
	if curParent == nil {
		return cur
	}
	if curParent.Left != nil && curParent.Left.Val == key {
		curParent.Left = cur
	} else {
		curParent.Right = cur
	}
	return root
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
