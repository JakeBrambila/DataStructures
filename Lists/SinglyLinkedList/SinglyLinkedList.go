package main

import "fmt"

// node of a linked list
type Node struct {
	Value int
	Next  *Node
}

// linked list definition
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// adds a node to the beginning by having the node point to head and then setting the head as the new node
func (l *LinkedList) AddToBeginning(n *Node) {
	//empty list case
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
		return
	}
	n.Next = l.Head
	l.Head = n
	l.Length++
}

// adds a node to the end of the linked list
func (l *LinkedList) AddToEnd(n *Node) {

	//empty list case
	if l.Tail == nil {
		l.Head = n
		l.Tail = n
		l.Length++
		return
	}

	l.Tail.Next = n
	l.Tail = n
	l.Length++
}

// prints out the linked list
func (l *LinkedList) Print() {
	temp := l.Head
	for temp != l.Tail {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Next
	}
	fmt.Println(temp.Value)
}
