package main

//230. 二叉搜索树中第K小的元素
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder(root, k, &count, &res)
	return res
}

//中序遍历
func inorder(node *TreeNode, k int, count *int, res *int) {
	if node != nil {
		inorder(node.Left, k, count, res)
		*count++
		if *count == k {
			*res = node.Val
			return
		}
		inorder(node.Right, k, count, res)
	}
}

func kthSmallest1(root *TreeNode, k int) int {
	arr := make([]int, 0)
	inOrder1(root, &arr)
	return arr[k-1]
}

func inOrder1(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder1(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder1(root.Right, arr)
}
