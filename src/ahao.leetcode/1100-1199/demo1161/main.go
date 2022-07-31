package main

/*
1161. 最大层内元素和
*/

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

//深度优先
func maxLevelSum(root *TreeNode) (ans int) {
	sum := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if level == len(sum) {
			sum = append(sum, node.Val)
		} else {
			sum[level] += node.Val
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	for i, s := range sum {
		if s > sum[ans] {
			ans = i
		}
	}
	return ans + 1 // 层号从 1 开始
}


//广度优先
func maxLevelSum2(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		q = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}
