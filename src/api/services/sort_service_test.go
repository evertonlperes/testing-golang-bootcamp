package services

import "testing"

func TestConstant(t *testing.T) {
	if privateConst != "private" {
		t.Error("The privateConst should be 'private'. Found", privateConst)
	}

	if publicConst != "public" {
		t.Error("The publicConst should be 'public'. Found", privateConst)
	}

}

func TestSort(t *testing.T) {
	elements := []int{9, 6, 5, 3, 8, 4, 7, 1, 2, 3, 0}

	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0. Found", elements[0])
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9. Found", elements[len(elements)-1])
	}
}
