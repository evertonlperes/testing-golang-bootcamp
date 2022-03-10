package sort

import (
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init layer (optional) - create a slice of elements (integers)
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution layer - invoking the class under test
	BubbleSort(elements)

	// Validation(s)
	if elements[0] != 9 {
		t.Error("First element to be 9, but found", elements[0])
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0, but found", len(elements)-1)
	}
}
