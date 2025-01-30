package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedList) AddBeginning(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
	} else {
		l.Head.Prev = n
		n.Next = l.Head
		l.Head = n
		l.Length++
	}
}

func (l *LinkedList) PrintList() {
	temp := l.Head
	for temp != l.Tail {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Next
	}
	fmt.Printf("%d", temp.Value)
}

func (l *LinkedList) PrintListReverse() {
	temp := l.Tail
	for temp != l.Head {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Prev
	}
	fmt.Printf("%d", temp.Value)
}

func (l *LinkedList) AddToEnd(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
	}
	l.Tail.Next = n
	n.Prev = l.Tail
	l.Tail = n
	l.Length++
}
