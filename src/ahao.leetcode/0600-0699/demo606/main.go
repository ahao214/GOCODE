package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"strconv"
	"strings"
)

/*
606. 根据二叉树创建字符串
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}

//迭代
func tree2str2(root *TreeNode) string {
	ans := &strings.Builder{}
	st := []*TreeNode{root}
	vis := map[*TreeNode]bool{}
	for len(st) > 0 {
		node := st[len(st)-1]
		if vis[node] {
			if node != root {
				ans.WriteByte(')')
			}
			st = st[:len(st)-1]
		} else {
			vis[node] = true
			if node != root {
				ans.WriteByte('(')
			}
			ans.WriteString(strconv.Itoa(node.Val))
			if node.Left == nil && node.Right != nil {
				ans.WriteString("()")
			}
			if node.Right != nil {
				st = append(st, node.Right)
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
		}
	}
	return ans.String()
}
