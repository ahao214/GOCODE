package main

import (
	"strconv"
)

//257. 二叉树的所有路径
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTreePath(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	tmpLeft := findTreePath(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}

	tmpRight := findTreePath(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res
}

//方法二
func findTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	helper(root, "", &res)
	return res
}

func helper(root *TreeNode, path string, res *[]string) {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	path += "->"
	if root.Left != nil {
		helper(root.Left, path, res)
	}
	if root.Right != nil {
		helper(root.Right, path, res)
	}
}
