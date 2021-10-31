package main

//653. 两数之和 IV - 输入 BST
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
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
