package arrays

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		// I could perform the same check from below using `reflect.DeepEqual`, but
		// that is not type safe.
		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}

	}

	t.Run("sum the tails of n slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("safely sum tails including empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}
