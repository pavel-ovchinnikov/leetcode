package largestpositiveintegerthatexistswithitsnegative

import "testing"

func Test_findMaxK(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 2, -3, 3, 5},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{-1, 10, 6, 7, -7, 1},
			want: 7,
		},
		{
			name: "Example 3",
			nums: []int{-10, 8, 6, 7, -2, -3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxK(tt.nums)
			if got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}
