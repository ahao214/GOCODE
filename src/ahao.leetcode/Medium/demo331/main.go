package main

import (
	"fmt"
	"strings"
)

//题目：331. 验证二叉树的前序序列化
//说明：序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录
//输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
//输出: true

func isValidSerialization(preorder string) bool {
	nodes, diff := strings.Split(preorder, ","), 1
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}

func main() {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	res := isValidSerialization(str)
	fmt.Println(res)
}
