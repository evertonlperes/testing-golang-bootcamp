package services

import "testing-golang/src/api/utils/sort"

const (
	privateConst = "private"
	publicConst  = "public"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
