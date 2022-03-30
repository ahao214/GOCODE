package main

/*
653. 两数之和 IV - 输入 BST

难度：简单

给定一个二叉搜索树 root 和一个目标结果 k，如果 BST
中存在两个元素且它们的和等于给定的目标结果，则返回 true。



示例 1：


输入: root = [5,3,6,2,4,null,7], k = 9
输出: true
示例 2：


输入: root = [5,3,6,2,4,null,7], k = 28
输出: false


提示:

二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
root 为二叉搜索树
-105 <= k <= 105
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

func findTarget2(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	res := []int{}
	inorder(root, &res)
	left, right := 0, len(res)-1
	for left < right {
		sum := res[left] + res[right]
		if sum == k {
			return true
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return false
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
