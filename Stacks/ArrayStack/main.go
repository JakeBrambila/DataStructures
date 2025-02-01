package main

import (
	"fmt"
)

func main() {
	// Create a new stack
	stack := Stack{}

	// Test Push function
	fmt.Println("Pushing elements onto the stack...")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	// Test PrintStack function
	fmt.Println("Stack after pushing elements:")
	stack.PrintStack()
	fmt.Println()

	// Test Pop function
	fmt.Println("Popping elements from the stack...")
	for i := 0; i < 5; i++ {
		value, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Popped value: %d\n", value)
		}
	}

	// Test PrintStack function after popping elements
	fmt.Println("Stack after popping elements:")
	stack.PrintStack()
	fmt.Println()

	// Test Pop function on an empty stack
	fmt.Println("Attempting to pop from an empty stack...")
	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped value: %d\n", value)
	}
}
