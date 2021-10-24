package main

//515. 在每个树行中找最大值
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		length := len(q) //每层的长度
		level := []int{}
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		//find the max of each level
		maxValue := max(level)
		res = append(res, maxValue)
	}
	return res
}

func max(level []int) int {
	if len(level) == 0 {
		return -1
	}
	res := level[0]
	for _, v := range level {
		if v > res {
			res = v
		}
	}
	return res
}
