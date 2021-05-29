package search

func BinarySearch(list []int, item int) int {
	low := 0
	max := len(list) - 1

	for low <= max {
		mid := (max + low) / 2
		guess := list[mid]

		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else if guess > item {
			max = mid - 1
		}
	}
	return -1
}
