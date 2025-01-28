package main

import (
	"errors"
	"fmt"
)

// dynamic array that is built using a static array
// to intialize this you would need something like this:
// arr := DynamicArray[int]{Array: make([]int, 0)}
// you could also intialize the Len and Capacity variables as well if you wanted to
type DynamicArray[T comparable] struct {
	Array    []T
	Len      int //length the user thinks the array is
	Capacity int //length of the actual array
}

// returns a pointer to a dynamic array and an error if the
// capacity is less than 0
func NewDynamicArray[T comparable](capacity int) (*DynamicArray[T], error) {
	if capacity < 0 {
		return nil, errors.New("Invalid Capacity")
	}
	return &DynamicArray[T]{
		Array:    make([]T, capacity),
		Capacity: capacity,
	}, nil
}

// Getter for length of the array that the user can see
func (arr *DynamicArray[T]) GetLength() int {
	return arr.Len
}

// checks if the array is empty
func (arr *DynamicArray[T]) IsEmpty() bool {
	return arr.GetLength() == 0
}

// returns the static array
func (arr *DynamicArray[T]) GetArray() []T {
	return arr.Array
}

// sets an element at a given index
func (arr *DynamicArray[T]) set(index int, elem T) error {
	if index < 0 || index > arr.Capacity {
		return errors.New("Invalid Index")
	}
	arr.Array[index] = elem
	return nil
}

// function to clear all values from the underlying array
func (arr *DynamicArray[T]) clear() {
	//intializes zero to the "zero" value for the generic T
	var zero T

	for i := 0; i < arr.Capacity; i++ {
		arr.Array[i] = zero
	}
	arr.Len = 0
}

// adds an element to the end of the Dynamic Array
// if our static underlying array is out of capacity then we double the capacity
// and copy the old array into the new array that has double the capacity
func (arr *DynamicArray[T]) add(elem T) {

	//checks if we are at capacity
	// if we are then we double the size of the array
	if arr.Len >= arr.Capacity {
		if arr.Capacity == 0 {
			arr.Capacity = 1
		} else {
			arr.Capacity *= 2
		}
		newArr := make([]T, arr.Capacity)

		//copy old array into new array
		for i := 0; i < arr.Len; i++ {
			newArr[i] = arr.Array[i]
		}
		//set the new Array as the one being used in the dynamic Array
		arr.Array = newArr
	}

	//add the new element and increase the length
	arr.Array[arr.Len] = elem
	arr.Len++
}

// removes a specific value of an array at a certain index
func (arr *DynamicArray[T]) RemoveAt(rIndex int) error {
	if rIndex >= arr.Len || rIndex < 0 {
		return errors.New("Index is out of bounds")
	}

	//moves all items in the array to the left
	for i := rIndex; i < arr.Len-1; i++ {
		arr.Array[i] = arr.Array[i+1]
	}

	//subtracts the length
	arr.Len--

	var zero T
	arr.Array[arr.Len] = zero
	return nil
}

// returns the index if the value is in the array
// if not it return a -1
func (arr *DynamicArray[T]) indexOf(elem T) int {
	for i := 0; i < arr.Len; i++ {
		if arr.Array[i] == elem {
			return i
		}
	}
	return -1
}

// returns true if the array contains an element
// false if it doesn't
func (arr *DynamicArray[T]) contains(elem T) bool {
	return arr.indexOf(elem) != -1
}

// returns a string format of the array
func (arr *DynamicArray[T]) ToString() string {
	if arr.Len == 0 {
		return "[]"
	}
	var arrString = "["
	for i := 0; i < arr.Len; i++ {
		arrString += fmt.Sprintf("%v", arr.Array[i])

		if i < arr.Len-1 {
			arrString += ", "
		}
	}
	arrString += "]"
	return arrString
}
