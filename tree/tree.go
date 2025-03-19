package tree

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}
	if data <= node.Data {
		node.Left = Insert(node.Left, data)
	} else {
		node.Right = Insert(node.Right, data)
	}
	return node
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}
