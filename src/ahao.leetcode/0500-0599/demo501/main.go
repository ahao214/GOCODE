package main

//501. 二叉搜索树中的众数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	max := 0
	dic := make(map[int]int)
	inOrder(root, &dic)
	for _, v := range dic {
		if v > max {
			max = v
		}
	}
	for k, v := range dic {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func inOrder(root *TreeNode, dic *map[int]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, dic)
	(*dic)[root.Val]++
	inOrder(root.Right, dic)
}
