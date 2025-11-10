package numberofarithmetictriplets

import "testing"

func Test_arithmeticTriplets(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		diff int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 4, 6, 7, 10},
			diff: 3,
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{4, 5, 6, 7, 8, 9},
			diff: 2,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arithmeticTriplets(tt.nums, tt.diff)
			if got != tt.want {
				t.Errorf("arithmeticTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
