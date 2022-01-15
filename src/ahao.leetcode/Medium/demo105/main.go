package main

/*
105. 从前序与中序遍历序列构造二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	return help(preorder, inorder)
}

func help(pre []int, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	//根结点
	rootVal := pre[0]
	index := 0
	for i, v := range in {
		if v == rootVal {
			index = i
		}
	}
	root := &TreeNode{}
	root.Val = rootVal
	root.Left = help(pre[1:index+1], in[:index])
	root.Right = help(pre[index+1:], in[index+1:])
	return root
}
