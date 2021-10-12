package moduleone

func Fibonacci(n int) int {
	fibn := [2]int{0, 1}
	for i := 0; i < n; i++ {
		next := fibn[0] + fibn[1]
		fibn[0], fibn[1] = fibn[1], next
	}
	return fibn[0]
}
