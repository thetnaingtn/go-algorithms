package sort

func findSmallest(list []int) int {
	smallestItem := list[0]
	smallestIndex := 0

	for i, item := range list {
		if item < smallestItem {
			smallestIndex = i
			smallestItem = item
		}
	}

	return smallestIndex
}

func SelectionSort(list []int) []int {
	sortedList := make([]int, len(list))
	for i := range list {
		index := findSmallest(list)
		sortedList[i] = list[index]
		list = append(list[:index], list[index+1:]...)
	}

	return sortedList
}
