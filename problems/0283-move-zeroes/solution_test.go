package movezeroes

import "testing"

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
	}{
		{
			name: "example 1",
			nums: []int{0, 1, 0, 3, 12},
		},
		{
			name: "example 2",
			nums: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
		})
	}
}
