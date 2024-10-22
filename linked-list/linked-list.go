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

func (ll *LinkedList) InsertAtEnd(value string) {
	newNode := &Node{Data: value, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
