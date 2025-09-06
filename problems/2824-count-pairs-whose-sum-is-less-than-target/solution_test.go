package countpairswhosesumislessthantarget

import "testing"

func Test_countPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 1, 2, 3, 1},
			target: 2,
			want:   3,
		},
		{
			name:   "Example 2",
			nums:   []int{-6, 2, 5, -2, -7, -1, 3},
			target: -2,
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPairs(tt.nums, tt.target)

			if got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
