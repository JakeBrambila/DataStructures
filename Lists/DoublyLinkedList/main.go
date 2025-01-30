package main

import "fmt"

func main() {
	// Create a new linked list
	list := LinkedList{}

	// Test adding nodes to the beginning of the list
	node1 := &Node{Value: 10}
	node2 := &Node{Value: 20}
	node3 := &Node{Value: 30}

	list.AddBeginning(node1)
	list.AddBeginning(node2)
	list.AddBeginning(node3)

	// Print the list from head to tail
	fmt.Println("Linked List (Head to Tail):")
	list.PrintList()

	fmt.Println()

	// Print the list from tail to head
	fmt.Println("Linked List (Tail to Head):")
	list.PrintListReverse()

	fmt.Println()

	// Print the length of the list
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test adding nodes to the end of the list
	node4 := &Node{Value: 40}
	node5 := &Node{Value: 50}
	node6 := &Node{Value: 60}

	list.AddToEnd(node4)
	list.AddToEnd(node5)
	list.AddToEnd(node6)

	// Print the list from head to tail
	fmt.Println("Linked List after adding to the end (Head to Tail):")
	list.PrintList()

	fmt.Println()

	// Print the list from tail to head
	fmt.Println("Linked List after adding to the end (Tail to Head):")
	list.PrintListReverse()

	fmt.Println()

	// Print the length of the list
	fmt.Printf("Length of the list: %d\n", list.Length)
}
