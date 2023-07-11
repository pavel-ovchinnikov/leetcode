package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	weights  []int
	days     int
	expected int
}

var tests = []testCase{
	// {
	// 	[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	5,
	// 	15,
	// },
	// {
	// 	[]int{3, 2, 2, 4, 1, 4},
	// 	3,
	// 	6,
	// },
	{
		[]int{1, 2, 3, 1, 1},
		4,
		3,
	},
}

func Test(t *testing.T) {
	for _, test := range tests {
		if output := shipWithinDays(test.weights, test.days); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
			fmt.Printf("t: %v\n", t)
		}
	}
}
