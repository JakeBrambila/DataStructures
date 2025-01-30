package main

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

func (l *LinkedList) AddToEnd(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
	}
}
