package moduletwo

func BubbleSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		swap := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return list
}
