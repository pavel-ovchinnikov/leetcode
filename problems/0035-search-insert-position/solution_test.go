package searchinsertposition

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{name: "test 1", nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{name: "test 2", nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{name: "test 3", nums: []int{1, 3, 5, 6}, target: 7, want: 4},
		{name: "test 4", nums: []int{1}, target: 7, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
