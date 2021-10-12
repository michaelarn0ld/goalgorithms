package moduleone

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	type Test struct {
		word string
		want string
	}

	tests := []Test{
		{"cat", "tac"},
		{"alphabet", "tebahpla"},
		{"日本語", "語本日"},
		{"banana pie", "eip ananab"},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("%v", testcase.word),
			func(t *testing.T) {
				result := ReverseString(testcase.word)
				if result != testcase.want {
					t.Fatalf("ReverseString() = %v; want %v", result, testcase.want)
				}
			})
	}
}
