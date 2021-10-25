package main

//116. 填充每个节点的下一个右侧节点指针
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectNode(root.Left, root.Right)
	return root
}

func connectNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectNode(node1.Left, node1.Right)
	connectNode(node2.Left, node2.Right)
	connectNode(node1.Right, node2.Left)
}
