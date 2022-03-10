package services

import "testing-golang/src/api/utils/sort"

const (
	privateConst = "private"
	publicConst  = "public"
)

func Sort(elements []int) {
	sort.BubbleSort(elements)
}
