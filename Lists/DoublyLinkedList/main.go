package main

import "fmt"

func main() {
	// Create a new linked list
	list := LinkedList{}

	// Test 1: AddBeginning
	fmt.Println("Test 1: AddBeginning")
	node1 := &Node{Value: 10}
	list.AddBeginning(node1)
	fmt.Println("Linked List after adding 10 to the beginning:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 2: AddBeginning
	fmt.Println("\nTest 2: AddBeginning")
	node2 := &Node{Value: 20}
	list.AddBeginning(node2)
	fmt.Println("Linked List after adding 20 to the beginning:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 3: AddToEnd
	fmt.Println("\nTest 3: AddToEnd")
	node3 := &Node{Value: 30}
	list.AddToEnd(node3)
	fmt.Println("Linked List after adding 30 to the end:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 4: AddToEnd
	fmt.Println("\nTest 4: AddToEnd")
	node4 := &Node{Value: 40}
	list.AddToEnd(node4)
	fmt.Println("Linked List after adding 40 to the end:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 5: InsertByIndex
	fmt.Println("\nTest 5: InsertByIndex")
	node5 := &Node{Value: 25}
	list.InsertByIndex(2, node5)
	fmt.Println("Linked List after inserting 25 at index 2:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 6: InsertByIndex
	fmt.Println("\nTest 6: InsertByIndex")
	node6 := &Node{Value: 35}
	list.InsertByIndex(4, node6)
	fmt.Println("Linked List after inserting 35 at index 4:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 7: isInList
	fmt.Println("\nTest 7: isInList")
	fmt.Println("Is 25 in the list? Position:", list.isInList(25))   // Expected: 2
	fmt.Println("Is 100 in the list? Position:", list.isInList(100)) // Expected: -1

	// Test 8: isInList
	fmt.Println("\nTest 8: isInList")
	fmt.Println("Is 40 in the list? Position:", list.isInList(40)) // Expected: 5
	fmt.Println("Is 20 in the list? Position:", list.isInList(20)) // Expected: 1

	// Test 9: RemoveBeginning
	fmt.Println("\nTest 9: RemoveBeginning")
	list.RemoveBeginning()
	fmt.Println("Linked List after removing the beginning:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 10: RemoveBeginning
	fmt.Println("\nTest 10: RemoveBeginning")
	list.RemoveBeginning()
	fmt.Println("Linked List after removing the beginning:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 11: RemoveEnd
	fmt.Println("\nTest 11: RemoveEnd")
	list.RemoveEnd()
	fmt.Println("Linked List after removing the end:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 12: RemoveEnd
	fmt.Println("\nTest 12: RemoveEnd")
	list.RemoveEnd()
	fmt.Println("Linked List after removing the end:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 13: RemoveByValue
	fmt.Println("\nTest 13: RemoveByValue")
	list.RemoveByValue(25)
	fmt.Println("Linked List after removing 25:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)

	// Test 14: RemoveByValue
	fmt.Println("\nTest 14: RemoveByValue")
	list.RemoveByValue(35)
	fmt.Println("Linked List after removing 35:")
	list.PrintList()
	fmt.Println("Linked List in reverse:")
	list.PrintListReverse()
	fmt.Printf("Length of the list: %d\n", list.Length)
}
