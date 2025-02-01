package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	stackArray []int
	size       int
}

func (s *Stack) Push(value int) {
	s.stackArray = append(s.stackArray, value)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, errors.New("Stack is Empty")
	}
	temp := s.stackArray[s.size-1]
	s.stackArray = s.stackArray[:s.size-1]
	s.size--
	return temp, nil
}

func (s *Stack) PrintStack() {
	for i := 0; i < s.size; i++ {
		fmt.Printf("%d - >", s.stackArray[i])
	}
}
