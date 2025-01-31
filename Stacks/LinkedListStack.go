package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Size int
	Head *Node
}

func (s *Stack) push(n *Node) {
	//empty list case
	if s.Head == nil {
		s.Head = n
		s.Size++

		return
	}
	n.Next = s.Head
	s.Head = n
	s.Size++
}

func (s *Stack) pop() (*Node, error) {
	if s.Head == nil {
		return nil, errors.New("Stack is empty")
	}
	temp := s.Head
	s.Head = s.Head.Next
	s.Size--

	return temp, nil
}

func (s *Stack) PrintStack() {
	temp := s.Head
	for i := 0; i < s.Size; i++ {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Next
	}
}
