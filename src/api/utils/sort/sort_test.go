package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// unit test
func TestBubbleSortOrderIncreasingOrder(t *testing.T) {

	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])

}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "First element should be 0. Found", elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9. Found", elements[len(elements)-1])
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	elements := GetElements(100.000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
