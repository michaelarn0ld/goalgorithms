package moduleone

import (
	"strconv"
)

func DecToBase(dec, base int) string {
	var str string
	for dec > 0 {
		rem := dec % base
		var chr string
		switch {
		case rem == 15:
			chr = "F"
		case rem == 14:
			chr = "E"
		case rem == 13:
			chr = "D"
		case rem == 12:
			chr = "C"
		case rem == 11:
			chr = "B"
		case rem == 10:
			chr = "A"
		default:
			chr = strconv.Itoa(rem)
		}
		str = chr + str
		dec /= base
	}
	return str
}
