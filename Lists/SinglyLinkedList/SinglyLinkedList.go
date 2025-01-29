package main

import (
	"fmt"
)

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

func (l *LinkedList) InsertAt(index int, n *Node) {
	//add to beginning
	if index == 0 {
		l.AddToBeginning(n)
		return
	}

	//add to end
	if index == l.Length-1 {
		l.AddToEnd(n)
		return
	}

	//inserts at any other index in the list
	temp := l.Head
	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	n.Next = temp.Next
	temp.Next = n
	l.Length++
}

// removes a node from the end of a Linked List
func (l *LinkedList) RemoveEnd() {
	temp := l.Head
	for temp.Next != l.Tail {
		temp = temp.Next
	}
	l.Tail = temp
	l.Length--
}

// removes a node from the beginning of a Linked List
func (l *LinkedList) RemoveBeginning() {
	l.Head = l.Head.Next
	l.Length--
}

func (l *LinkedList) RemoveAt(index int) {
	//remove from beginning
	if index == 0 {
		l.RemoveBeginning()
		return
	}

	//remove from end
	if index == l.Length-1 {
		l.RemoveEnd()
		return
	}

	//removes from any other index
	temp := l.Head
	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	l.Length--
}

// checks if a value is in a linked list
func (l *LinkedList) IsInList(value int) bool {
	temp := l.Head
	for temp != l.Tail {
		if temp.Value == value {
			return true
		}
		temp = temp.Next
	}

	//check tail
	if l.Tail.Value == value {
		return true
	}
	return false
}

// removes a node in the list by the value
func (l *LinkedList) RemoveValue(value int) {
	if l.IsInList(value) {

		//check if the value is the first Node
		if l.Head.Value == value {
			l.RemoveBeginning()
			l.Length--
		}

		temp := l.Head
		for temp.Next.Value != value {
			temp = temp.Next
		}
		temp.Next = temp.Next.Next
		l.Length--
	} else {
		fmt.Println("Value not in List.")
	}
}
