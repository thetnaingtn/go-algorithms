package recursion

func BinarySearch(list []int, item int) func() int {
	low := 0
	high := len(list) - 1

	//use closure to maintain low/high state
	var binary func() int
	binary = func() int {
		mid := (low + high) / 2
		predict := list[mid]
		if low <= high {
			if predict == item {
				return mid
			} else if predict < item {
				low = mid + 1
				return binary()
			} else {
				high = mid - 1
				return binary()
			}
		}

		return -1
	}

	return binary
}
