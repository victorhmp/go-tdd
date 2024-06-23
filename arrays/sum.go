package main

// Important to notice that arrays in Go have fixed length, and that length is
// fixed to its type declaration. So a function that expects [5]int will NOT
// accept an array that is not of length 5.

func Sum(numbers []int) int {
	sum := 0

	// Notice that range is returning (index, value)
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// From the docs:
// Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation [...]
//
// There are major differences between the ways arrays work in Go and C. In Go,
//
// - Arrays are values. Assigning one array to another copies all the elements.
// - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
// - The size of an array is part of its type. The types [10]int and [20]int are distinct.

// About Slices:
// Slices wrap arrays to give a more general, powerful, and convenient interface
// to sequences of data. Except for items with explicit dimension such as
// transformation matrices, most array programming in Go is done with slices
// rather than simple arrays.
// Important to note that slices are **references** to an underlying array.
