package sort

import (
	"math/rand"
	"time"
)

func QuickSort(list []int) []int {
	rand.Seed(time.Now().UnixNano())
	if len(list) < 2 {
		return list
	} else {
		var left, right, result []int
		pivotIndex := rand.Intn(len(list) - 1)
		pivot := list[pivotIndex]
		for _, itm := range list {
			if itm < pivot {
				left = append(left, itm)
			} else if itm > pivot {
				right = append(right, itm)
			} else {
				continue
			}
		}
		result = append(result, QuickSort(left)...)
		result = append(result, pivot)
		result = append(result, QuickSort(right)...)

		return result
	}
}
