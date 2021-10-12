package moduleone

// Sum sums all the numbers in a slice and returns the result
func Sum(list []int) int {
	result := 0
	for _, num := range list {
		result += num
	}
	return result
}
