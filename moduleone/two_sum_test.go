package moduleone

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type Test struct {
		numbers []int
		sum     int
		want    [2]int
	}

	tests := []Test{
		{[]int{1, 2, 3, 4}, 7, [2]int{2, 3}},
		{[]int{0, -1, 1}, 0, [2]int{1, 2}},
		{[]int{0, 1, 1}, 0, [2]int{-1, -1}},
		{[]int{10, 1, 12, 3, 7, 2, 2, 1}, 4, [2]int{1, 3}},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("%v with sum %v", testcase.numbers, testcase.sum), func(t *testing.T) {
			result := TwoSum(testcase.numbers, testcase.sum)
			if ok := compareArray(result, testcase.want); !ok {
				t.Fatalf("TwoSum() = %v; want %v", result, testcase.want)
			}
		})
	}
}

func compareArray(result [2]int, want [2]int) bool {
	for i, _ := range result {
		if result[i] != want[i] {
			return false
		}
	}
	return true
}
