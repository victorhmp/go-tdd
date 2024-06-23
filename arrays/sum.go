package main

// Important to notice that arrays in Go have fixed length, and that length is
// fixed to its type declaration. So a function that expects [5]int will NOT
// accept an array that is not of length 5.

func Sum(numbers [5]int) int {
	sum := 0

	// Notice that range is returning (index, value)
	for _, number := range numbers {
		sum += number
	}

	return sum
}
