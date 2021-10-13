package moduletwo

func InsertionSort(list []int) []int {
	var sorted []int
	for _, val := range list {
		sorted = addToList(val, sorted)
	}
	return sorted
}

func addToList(val int, sorted []int) []int {
	// nothing is in sorted, just put the number in
	if len(sorted) == 0 {
		return append(sorted, val)
	}
	// the sorted list is non-empty, so we need to know where to put stuff
	left := 0
	right := len(sorted) - 1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case val >= sorted[mid]:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	// put val at index left
	switch {
	case left >= len(sorted):
		sorted = append(sorted, val)
	case left < 0:
		sorted = append([]int{val}, sorted...)
	default:
		sorted = append(sorted[:left+1], sorted[left:]...)
		sorted[left] = val
	}
	return sorted
}
