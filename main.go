package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	list := []int{7, 4, 9, 18, 1, 1}
	sorted := sort.QuickSort(list)

	fmt.Println(sorted)
}
