package main

/*
106. 从中序与后序遍历序列构造二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}
	return help(inorder, postorder)
}

func help(in []int, post []int) *TreeNode {
	if len(in) == 0 || len(post) == 0 {
		return nil
	}
	rootVal := post[len(post)-1]
	index := 0
	for i, v := range in {
		if v == rootVal {
			index = i
		}
	}
	root := &TreeNode{}
	root.Val = rootVal
	root.Left = help(in[:index], post[:index])
	root.Right = help(in[index+1:], post[index:len(post)-1])
	return root
}
