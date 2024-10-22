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

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Print("null")
}

func (ll *LinkedList) Count() {
	current := ll.Head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	fmt.Println(count)
}

func (ll *LinkedList) RemoveFirstNode() {
	if ll.Head == nil {
		return
	}
	ll.Head = ll.Head.Next
}

func (ll *LinkedList) RemoveEndNode() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head

	for previous.Next.Next != nil {
		previous = previous.Next
	}
	previous.Next = nil
}