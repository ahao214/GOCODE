package main

/*
889. 根据前序和后序遍历构造二叉树
返回与给定的前序和后序遍历匹配的任何二叉树。
pre 和 post 遍历中的值是不同的正整数。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return help(preorder, postorder)
}

func help(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 {
		return nil
	}
	if len(pre) == 1 || len(post) == 1 {
		root := &TreeNode{}
		root.Val = pre[0]
		return root
	}
	index := 0
	for i, v := range post {
		if v == pre[1] {
			index = i + 1
		}
	}
	root := &TreeNode{}
	root.Val = pre[0]
	root.Left = help(pre[1:index+1], post[:index])
	root.Right = help(pre[index+1:], post[index:len(post)-1])
	return root
}
