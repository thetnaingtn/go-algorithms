package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	// list := []int{1, 2, 3, 4, 5}
	// item := 7
	// found := search.BinarySearch(list, item)
	// fmt.Printf("Found %d at position %d\n", item, found)
	list := []int{1, 9, 10, 5, 8, 3}
	sorted := sort.SelectionSort(list)

	fmt.Println(sorted)
}
