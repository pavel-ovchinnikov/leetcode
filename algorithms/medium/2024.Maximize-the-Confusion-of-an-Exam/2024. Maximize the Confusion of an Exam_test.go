package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	answerKey string
	k         int
	expected  int
}

var tests = []testCase{
	{
		"TTFF",
		2,
		4,
	},
	{
		"TFFT",
		1,
		3,
	},
	{
		"TTFTTFTT",
		1,
		5,
	},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := maxConsecutiveAnswers(test.answerKey, test.k); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
			fmt.Printf("t: %v\n", t)
		}
	}
}
