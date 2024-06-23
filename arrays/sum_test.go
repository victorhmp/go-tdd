package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Empty collection", func(t *testing.T) {
		numbers := []int{}

		got := Sum(numbers)
		expected := 0

		if got != expected {
			t.Errorf("got %d expected %d input: %v", got, expected, numbers)
		}
	})

	t.Run("Collection of n items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d input: %v", got, expected, numbers)
		}
	})
}
