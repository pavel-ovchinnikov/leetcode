package leetcode

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

var tests = []testCase{
	{[]int{1, 2, 5, 2, 3}, 2, []int{1, 2}},
	{[]int{1, 2, 5, 2, 3}, 3, []int{3}},
	{[]int{1, 2, 5, 2, 3}, 5, []int{4}},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := targetIndices(test.nums, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
			// slices.Equal(output, test.expected)
		}
	}
}
