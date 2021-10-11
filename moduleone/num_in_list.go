package moduleone

// NumInList returns true if a number is in the list otherwise it returns false
func NumInList(list []int, num int) bool {
	for _, item := range list {
		if item == num {
			return true
		}
	}
	return false
}
