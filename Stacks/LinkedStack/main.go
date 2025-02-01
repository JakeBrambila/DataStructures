package main

import (
	"fmt"
)

func main() {
	// Create a new stack
	stack := &Stack{}

	// Push some nodes onto the stack
	stack.push(&Node{Value: 10})
	stack.push(&Node{Value: 20})
	stack.push(&Node{Value: 30})

	// Print the stack
	fmt.Println("Stack after pushing 10, 20, 30:")
	stack.PrintStack()
	fmt.Println()

	// Pop a node from the stack
	poppedNode, err := stack.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped node value: %d\n", poppedNode.Value)
	}

	// Print the stack after popping
	fmt.Println("Stack after popping one element:")
	stack.PrintStack()
	fmt.Println()

	// Pop another node from the stack
	poppedNode, err = stack.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped node value: %d\n", poppedNode.Value)
	}

	// Print the stack after popping again
	fmt.Println("Stack after popping another element:")
	stack.PrintStack()
	fmt.Println()

	// Pop the last node from the stack
	poppedNode, err = stack.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped node value: %d\n", poppedNode.Value)
	}

	// Print the stack after popping the last element
	fmt.Println("Stack after popping the last element:")
	stack.PrintStack()
	fmt.Println()

	// Try to pop from an empty stack
	poppedNode, err = stack.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped node value: %d\n", poppedNode.Value)
	}
}
