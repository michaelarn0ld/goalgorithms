package moduleone

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("Fizz Buzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		case i%5 == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Printf("%v", i)
		}
		if i < n {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}
