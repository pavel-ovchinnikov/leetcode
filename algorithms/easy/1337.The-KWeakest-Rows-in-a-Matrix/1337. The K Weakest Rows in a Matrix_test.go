package leetcode

import (
	"reflect"
	"testing"
)

type testCase struct {
	mat      [][]int
	k        int
	expected []int
}

var tests = []testCase{
	{
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		3,
		[]int{2, 0, 3},
	},
	{
		[][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
		2,
		[]int{0, 2},
	},
	{
		[][]int{
			{1, 1, 0},
			{1, 1, 0},
			{1, 1, 1},
			{1, 1, 1},
			{0, 0, 0},
			{1, 1, 1},
			{1, 0, 0},
		},
		6,
		[]int{4, 6, 0, 1, 2, 3},
	},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := kWeakestRows(test.mat, test.k); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
