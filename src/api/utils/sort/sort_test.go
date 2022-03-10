package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortOrder(t *testing.T) {
	// Init layer (optional) - create a slice of elements (integers)
	elements := []int{7, 9, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println(elements)
	// Execution layer - invoking the class under test
	BubbleSort(elements)

	// Validation(s)
	if elements[0] != 0 {
		fmt.Println(elements)
		t.Error("First element to be 0, but found", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9, but found", len(elements)-1)
	}
}

func TestSortIncreasingOrder(t *testing.T) {
	// Init layer (optional) - create a slice of elements (integers)
	elements := []int{7, 9, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution layer - invoking the class under test
	Sort(elements)

	// Validation(s)
	if elements[0] != 0 {
		t.Error("First element to be 0, but found", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9, but found", len(elements)-1)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{7, 9, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{7, 9, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
