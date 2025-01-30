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

	// Test inserting nodes at specific indices
	node7 := &Node{Value: 15}
	node8 := &Node{Value: 25}
	node9 := &Node{Value: 55}

	// Insert at the beginning
	list.InsertByIndex(0, node7)
	fmt.Println("Linked List after inserting 15 at index 0 (Head to Tail):")
	list.PrintList()

	fmt.Println()

	// Insert in the middle
	list.InsertByIndex(3, node8)
	fmt.Println("Linked List after inserting 25 at index 3 (Head to Tail):")
	list.PrintList()

	fmt.Println()

	// Insert at the end
	list.InsertByIndex(list.Length-1, node9)
	fmt.Println("Linked List after inserting 55 at the end (Head to Tail):")
	list.PrintList()

	fmt.Println()

	// Print the list from tail to head
	fmt.Println("Linked List after all insertions (Tail to Head):")
	list.PrintListReverse()

	fmt.Println()

	// Print the length of the list
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Additional Test Case 1: Insert at an invalid index (negative index)
	node10 := &Node{Value: 100}
	fmt.Println("Attempting to insert 100 at index -1:")
	list.InsertByIndex(-1, node10)
	list.PrintList()

	fmt.Println()

	// Additional Test Case 2: Insert at an invalid index (index greater than the length of the list)
	node11 := &Node{Value: 200}
	fmt.Println("Attempting to insert 200 at index 20:")
	list.InsertByIndex(20, node11)
	list.PrintList()

	fmt.Println()

	// Additional Test Case 3: Insert at index 1 (closer to the head)
	node12 := &Node{Value: 5}
	fmt.Println("Attempting to insert 5 at index 1:")
	list.InsertByIndex(1, node12)
	list.PrintList()

	fmt.Println()

	// Additional Test Case 4: Insert at index 3 (closer to the head)
	node13 := &Node{Value: 35}
	fmt.Println("Attempting to insert 35 at index 3:")
	list.InsertByIndex(3, node13)
	list.PrintList()

	fmt.Println()

	// Additional Test Case 5: Insert at index 5 (closer to the head)
	node14 := &Node{Value: 45}
	fmt.Println("Attempting to insert 45 at index 5:")
	list.InsertByIndex(5, node14)
	list.PrintList()

	fmt.Println()
}
