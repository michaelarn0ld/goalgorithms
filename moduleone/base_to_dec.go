package moduleone

import (
	"math"
	"strconv"
)

func BaseToDec(str string, base int) int {
	var result int
	for i, j := len(str)-1, 0; i >= 0; i, j = i-1, j+1 {
		val := string(str[i])
		var num int
		switch {
		case val == "F":
			num = 15
		case val == "E":
			num = 14
		case val == "D":
			num = 13
		case val == "C":
			num = 12
		case val == "B":
			num = 11
		case val == "A":
			num = 10
		default:
			num, _ = strconv.Atoi(val)
		}
		result += num * int(math.Pow(float64(base), float64(j)))
	}
	return result
}
