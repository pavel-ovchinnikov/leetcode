package numberofgoodpairs

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1, 1, 3},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 1, 1},
			want: 6,
		},
		{
			name: "Example 3",
			nums: []int{1, 2, 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIdenticalPairs(tt.nums)
			if got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
