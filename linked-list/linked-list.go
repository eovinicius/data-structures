package linkedlist

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data any
	Next *Node
}

func (n *Node) Print() {
	current := n
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Print("null")
}
