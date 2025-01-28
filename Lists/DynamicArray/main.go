package main

import (
	"fmt"
)

func main() {
	// Test NewDynamicArray
	arr, err := NewDynamicArray[int](5)
	if err != nil {
		fmt.Println("Error creating array:", err)
		return
	}

	// Test IsEmpty
	fmt.Println("Is array empty?", arr.IsEmpty()) // true

	// Test Add
	arr.add(10)
	arr.add(20)
	arr.add(30)
	arr.add(40)
	arr.add(50)
	fmt.Println("Array after adding elements:", arr.ToString()) // [10, 20, 30, 40, 50]

	// Test GetLength
	fmt.Println("Length of array:", arr.GetLength()) // 5

	// Test GetArray
	fmt.Println("Underlying array:", arr.GetArray())

	// Test Set
	err = arr.set(2, 100)
	if err != nil {
		fmt.Println("Error setting value:", err)
	} else {
		fmt.Println("Array after setting index 2 to 100:", arr.ToString()) // [10, 20, 100, 40, 50]
	}

	// Test IndexOf
	index := arr.indexOf(40)
	fmt.Println("Index of 40:", index) // 3

	// Test Contains
	fmt.Println("Does array contain 20?", arr.contains(20)) // true
	fmt.Println("Does array contain 99?", arr.contains(99)) // false

	// Test RemoveAt
	err = arr.RemoveAt(1)
	if err != nil {
		fmt.Println("Error removing element:", err)
	} else {
		fmt.Println("Array after removing element at index 1:", arr.ToString()) // [10, 100, 40, 50]
	}

	// Test Clear
	arr.clear()
	fmt.Println("Array after clearing:", arr.ToString()) // []

	// Test dynamic resizing by adding elements beyond initial capacity
	for i := 1; i <= 10; i++ {
		arr.add(i * 10)
	}
	fmt.Println("Array after adding 10 elements:", arr.ToString()) // [10, 20, 30, ..., 100]
	fmt.Println("Length:", arr.GetLength(), "Capacity:", arr.Capacity)
}
