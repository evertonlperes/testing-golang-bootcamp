package sort

import (
	"fmt"
	"testing"
)

// unit test
func TestBubbleSortOrder(t *testing.T) {
	// Init layer (optional) - create a slice of elements (integers)
	elements := GetElements(10)
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
	elements := GetElements(10)

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
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
