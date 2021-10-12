package moduleone

func TwoSum(numbers []int, sum int) [2]int {
	m := make(map[int]int)
	for i, num := range numbers {
		if val, ok := m[sum-num]; ok {
			return [2]int{val, i}
		}
		m[num] = i
	}
	return [2]int{-1, -1}
}
