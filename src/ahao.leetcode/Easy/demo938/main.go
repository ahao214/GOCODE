package main

//938. 二叉搜索树的范围和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	arr := []int{}
	inOrder(root, &arr)
	res := 0
	for _, v := range arr {
		if v >= low && v <= high {
			res += v
		}
	}
	return res
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

//递归方法

func rangeSumBST1(root *TreeNode, low int, high int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val > low && root.Val > high {
		return rangeSumBST1(root.Left, low, high)
	} else if root.Val < low && root.Val < high {
		return rangeSumBST1(root.Right, low, high)
	} else if root.Val >= low && root.Val <= high {
		res += root.Val
		res += rangeSumBST1(root.Right, low, high)
		res += rangeSumBST1(root.Left, low, high)
	}
	return res
}
