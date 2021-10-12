package moduleone

import (
	"fmt"
	"testing"
)

func TestBaseToBase(t *testing.T) {
	type Test struct {
		str     string
		base    int
		newBase int
		want    string
	}

	tests := []Test{
		{
			"E", 16, 2,
			"1110",
		}, {
			"11011110101011011011111011101111", 2, 3,
			"100122100210211112102",
		}, {
			"3735928559", 10, 4,
			"3132223123323233",
		}, {
			"8831A383B", 12, 16,
			"DEADBEEF",
		},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("%v in base %v to base %v", testcase.str, testcase.base, testcase.newBase), func(t *testing.T) {
			got := BaseToBase(testcase.str, testcase.base, testcase.newBase)
			if got != testcase.want {
				t.Fatalf("BaseToBase() = %v; want %v", got, testcase.want)
			}
		})
	}
}
