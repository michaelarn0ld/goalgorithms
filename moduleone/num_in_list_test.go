package moduleone

import (
	"fmt"
	"testing"
)

func TestNumInList(t *testing.T) {
	type Test struct {
		list []int
		num  int
		want bool
	}

	tests := []Test{
		{nil, 1, false},
		{[]int{}, 1, false},
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2, 3, 4, -1}, -1, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 8, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 2, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 5, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 4, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 1, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 8, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 9, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 3, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 0, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 88, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 23, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 44, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 123, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 321, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 32, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 7, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 6, false},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, -1, true},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, 1, false},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("%v %v", testcase.list, testcase.num),
			func(t *testing.T) {
				result := NumInList(testcase.list, testcase.num)
				if result != testcase.want {
					t.Fatalf("NumInList() = %v; want %v", result, testcase.want)
				}
			})
	}
}
