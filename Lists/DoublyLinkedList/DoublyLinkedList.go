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
	//empty list case
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
	fmt.Println()
}

func (l *LinkedList) PrintListReverse() {
	temp := l.Tail
	for temp != l.Head {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Prev
	}
	fmt.Printf("%d", temp.Value)
	fmt.Println()
}

func (l *LinkedList) AddToEnd(n *Node) {
	//empty list case
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

// function to insert to the list by index
func (l *LinkedList) InsertByIndex(index int, n *Node) {
	//beginning case
	if index == 0 {
		l.AddBeginning(n)
		return
	}

	//end case
	if index >= l.Length {
		l.AddToEnd(n)
		return
	}

	// invalid index case
	if index < 0 {
		fmt.Println("Invalid Index.")
		return
	}

	// attempting to optimize
	// seeing if the index lands closer to the head or Tail
	if index < (l.Length-1)/2 {
		temp := l.Head
		for i := 0; i < index-1; i++ {
			temp = temp.Next
		}
		temp.Next.Prev = n
		n.Next = temp.Next
		temp.Next = n
		n.Prev = temp
	} else {
		temp := l.Tail
		for i := l.Length - 1; i > index; i-- {
			temp = temp.Prev
		}
		temp.Prev.Next = n
		n.Prev = temp.Prev
		temp.Prev = n
		n.Next = temp
	}
	l.Length++
}

// function that checks if a value is in the Linked List
func (l *LinkedList) isInList(value int) int {
	temp := l.Head
	for i := 0; i < l.Length; i++ {
		if temp.Value == value {
			return i
		}
		temp = temp.Next
	}
	return -1
}

// function to remove a Node from the beginning
func (l *LinkedList) RemoveBeginning() {
	temp := l.Head.Next
	temp.Prev = nil
	l.Head = temp
	l.Length--
}

// function to remove from the end
func (l *LinkedList) RemoveEnd() {
	temp := l.Tail.Prev
	temp.Next = nil
	l.Tail = temp
	l.Length--
}

// function to remove a node by value
func (l *LinkedList) RemoveByValue(value int) {
	// checks if the value is even in the list
	pos := l.isInList(value)
	if pos != -1 {
		//check if the value is in the beginning
		if pos == 0 {
			l.RemoveBeginning()
			return
		}

		//check if the value is at the end
		if pos == l.Length-1 {
			l.RemoveEnd()
			return
		}

		// optimization to start at head or Tail
		if pos < (l.Length-1)/2 {
			temp := l.Head
			for i := 0; i < pos-1; i++ {
				temp = temp.Next
			}
			temp.Next.Next.Prev = temp
			temp.Next = temp.Next.Next
			l.Length--
		} else {
			temp := l.Tail
			for i := l.Length - 1; i > pos; i-- {
				temp = temp.Prev
			}
			temp.Prev.Prev.Next = temp
			temp.Prev = temp.Prev.Prev
			l.Length--
		}

	}
	fmt.Println("Value not in List.")
}
