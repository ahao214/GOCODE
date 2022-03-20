package main

//108. 将有序数组转换为二叉搜索树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return help(nums, 0, len(nums)-1)
}

func help(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: arr[mid]}
	root.Left = help(arr, left, mid-1)
	root.Right = help(arr, mid+1, right)
	return root
}

func sortedArrayToBST1(arr []int) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		root := &TreeNode{}
		root.Val = arr[0]
		return root
	}
	mid := len(arr) / 2
	root := &TreeNode{}
	root.Val = arr[mid]
	root.Left = sortedArrayToBST1(arr[:mid])
	root.Right = sortedArrayToBST1(arr[mid+1:])
	return root
}
