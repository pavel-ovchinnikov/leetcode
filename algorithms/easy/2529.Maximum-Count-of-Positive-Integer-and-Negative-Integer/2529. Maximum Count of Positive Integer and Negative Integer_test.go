package leetcode

import (
	"testing"
)

type testCase struct {
	nums     []int
	expected int
}

var tests = []testCase{
	{[]int{-2, -1, -1, 1, 2, 3}, 3},
	{[]int{-3, -2, -1, 0, 0, 1, 2}, 3},
	{[]int{5, 20, 66, 1314}, 4},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := maximumCount(test.nums); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
