package leetcode

import (
	"testing"
)

type testCase struct {
	grid     [][]int
	expected int
}

var tests = []testCase{
	{[][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3}},
		8},
	{[][]int{
		{3, 2},
		{1, 0}},
		0},
	{[][]int{
		{5, 1, 0},
		{-5, -5, -5}},
		3},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := countNegatives(test.grid); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
