package main

import "fmt"

func main() {
	// Create a new linked list
	ll := LinkedList{}

	// Test AddToBeginning
	fmt.Println("Testing AddToBeginning:")
	n1 := &Node{Value: 10}
	n2 := &Node{Value: 20}
	ll.AddToBeginning(n1)
	ll.AddToBeginning(n2)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 2

	// Test AddToEnd
	fmt.Println("\nTesting AddToEnd:")
	n3 := &Node{Value: 30}
	n4 := &Node{Value: 40}
	ll.AddToEnd(n3)
	ll.AddToEnd(n4)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 4

	// Test InsertAt
	fmt.Println("\nTesting InsertAt:")
	n5 := &Node{Value: 25}
	ll.InsertAt(2, n5) // Inserting 25 at index 2
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 5

	// Test InsertAt beginning
	fmt.Println("\nTesting InsertAt (beginning):")
	n6 := &Node{Value: 5}
	ll.InsertAt(0, n6)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 6

	// Test InsertAt end
	fmt.Println("\nTesting InsertAt (end):")
	n7 := &Node{Value: 50}
	ll.InsertAt(ll.Length-1, n7)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 7

	// Test RemoveAt
	fmt.Println("\nTesting RemoveAt:")
	ll.RemoveAt(2) // Removing the node at index 2
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 6

	// Test RemoveAt beginning
	fmt.Println("\nTesting RemoveAt (beginning):")
	ll.RemoveAt(0)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 5

	// Test RemoveAt end
	fmt.Println("\nTesting RemoveAt (end):")
	ll.RemoveAt(ll.Length - 1)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 4

	// Test IsInList
	fmt.Println("\nTesting IsInList:")
	fmt.Println(ll.IsInList(25))  // Expected Output: true
	fmt.Println(ll.IsInList(100)) // Expected Output: false

	// Test RemoveValue
	fmt.Println("\nTesting RemoveValue:")
	ll.RemoveValue(25)
	ll.Print()
	fmt.Println("Length:", ll.Length) // Expected Output: 3

	// Test RemoveValue for non-existent value
	fmt.Println("\nTesting RemoveValue (non-existent value):")
	ll.RemoveValue(100)               // Expected Output: "Value is not in List."
	fmt.Println("Length:", ll.Length) // Expected Output: 3
}
