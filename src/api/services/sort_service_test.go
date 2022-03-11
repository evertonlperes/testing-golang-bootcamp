package services

import (
	"testing"
	"testing-golang/src/api/utils/sort"
)

// Integration test
func TestConstant(t *testing.T) {
	if privateConst != "private" {
		t.Error("The privateConst should be 'private'. Found", privateConst)
	}

	if publicConst != "public" {
		t.Error("The publicConst should be 'public'. Found", privateConst)
	}

}

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0. Found", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9. Found", elements[len(elements)-1])
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0. Found", elements[0])
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("Last element should be 10000. Found", elements[len(elements)-1])
	}
}
