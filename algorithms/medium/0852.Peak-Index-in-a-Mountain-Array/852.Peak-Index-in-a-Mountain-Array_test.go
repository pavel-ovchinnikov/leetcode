package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	data     []int
	expected int
}

var tests = []testCase{
	{
		[]int{0, 1, 0},
		1,
	},
	{
		[]int{0, 1, 2, 0},
		2,
	},
	{
		[]int{0, 1, 10, 11, 12, 1},
		4,
	},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := peakIndexInMountainArray(test.data); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
			fmt.Printf("t: %v\n", t)
		}
	}
}
